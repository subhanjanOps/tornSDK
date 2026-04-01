from __future__ import annotations

import json
import re
from collections import defaultdict
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
SPEC_PATH = ROOT / "openapi.json"

SKIP_PATHS = {
    "/user/basic",
    "/user/profile",
    "/user/bars",
    "/user/battlestats",
    "/faction/basic",
    "/faction/{id}/basic",
}

PAYLOAD = '{"ok":true}'
QUERY_CODE = 'url.Values{"comment": {"sdk"}}'


def exported(name: str) -> str:
    return name[:1].upper() + name[1:]


def path_params(path: str) -> list[str]:
    return re.findall(r"{([^}]+)}", path)


def go_path_expr(path: str, params: list[str]) -> tuple[str, bool]:
    trimmed = path.lstrip("/")
    if not params:
        return f'"{trimmed}"', False

    fmt_path = re.sub(r"{[^}]+}", "%s", trimmed)
    args = ", ".join(params)
    return f'fmt.Sprintf("{fmt_path}", {args})', True


def method_signature(params: list[str]) -> str:
    parts = [f"{param} string" for param in params]
    parts.append("query url.Values")
    return ", ".join(parts)


def test_args(params: list[str]) -> str:
    if not params:
        return "query"

    samples = ", ".join(f'"{param}-value"' for param in params)
    return f"{samples}, query"


def expected_path(path: str, params: list[str]) -> str:
    result = path.lstrip("/")
    for param in params:
        result = result.replace("{" + param + "}", f"{param}-value")
    return result


def method_code(method_name: str, path: str, params: list[str]) -> str:
    signature = method_signature(params)
    path_expr, needs_fmt = go_path_expr(path, params)
    return (
        f"func (s *Service) {method_name}(ctx context.Context, {signature}) (rawapi.Response, error) {{\n"
        f"\treturn rawapi.Get(ctx, s.client, {path_expr}, query)\n"
        f"}}\n"
    )


def package_service_code(package: str, methods: list[dict], include_constructor: bool) -> str:
    needs_fmt = any(path_params(item["path"]) for item in methods)
    imports = ["\t\"context\""]
    if needs_fmt:
        imports.append('\t"fmt"')
    imports.extend(
        [
            '\t"net/url"',
            "",
            '\t"github.com/subhanjanOps/tornSDK/internal/rawapi"',
        ]
    )

    lines = [f"package {package}", "", "import (", *imports, ")", ""]

    if include_constructor:
        lines.extend(
            [
                "type Service struct {",
                "\tclient rawapi.Requester",
                "}",
                "",
                "func NewService(c rawapi.Requester) *Service {",
                "\treturn &Service{client: c}",
                "}",
                "",
            ]
        )

    for item in methods:
        lines.append(method_code(item["method_name"], item["path"], item["params"]).rstrip())
        lines.append("")

    return "\n".join(lines).rstrip() + "\n"


def package_test_code(package: str, methods: list[dict], use_existing_constructor: bool) -> str:
    constructor_call = "NewService" if use_existing_constructor else "NewService"
    lines = [
        f"package {package}",
        "",
        "import (",
        '\t"context"',
        '\t"encoding/json"',
        '\t"net/http"',
        '\t"net/url"',
        '\t"testing"',
        "",
        '\t"github.com/subhanjanOps/tornSDK/internal/httpclient"',
        '\t"github.com/subhanjanOps/tornSDK/internal/rawapi"',
        ")",
        "",
        "type rawStubRequester struct {",
        "\tt         *testing.T",
        "\twantPath  string",
        "\twantQuery url.Values",
        "\tpayload   string",
        "}",
        "",
        "func (s rawStubRequester) Do(_ context.Context, req *httpclient.Request, v interface{}) error {",
        "\ts.t.Helper()",
        "",
        '\tif got, want := req.Method, http.MethodGet; got != want {',
        '\t\ts.t.Fatalf("unexpected method: got %q want %q", got, want)',
        "\t}",
        "",
        '\tif got, want := req.Path, s.wantPath; got != want {',
        '\t\ts.t.Fatalf("unexpected path: got %q want %q", got, want)',
        "\t}",
        "",
        '\tif got, want := req.Query.Encode(), s.wantQuery.Encode(); got != want {',
        '\t\ts.t.Fatalf("unexpected query: got %q want %q", got, want)',
        "\t}",
        "",
        "\treturn json.Unmarshal([]byte(s.payload), v)",
        "}",
        "",
        f"func Test{exported(package)}RawMethods(t *testing.T) {{",
        f"\tquery := {QUERY_CODE}",
        "",
        "\tcases := []struct {",
        "\t\tname     string",
        "\t\twantPath string",
        "\t\tcall     func(*Service) (rawapi.Response, error)",
        "\t}{",
    ]

    for item in methods:
        args = test_args(item["params"])
        want_path = expected_path(item["path"], item["params"])
        lines.extend(
            [
                "\t\t{",
                f'\t\t\tname: "{item["method_name"]}",',
                f'\t\t\twantPath: "{want_path}",',
                f"\t\t\tcall: func(s *Service) (rawapi.Response, error) {{ return s.{item['method_name']}(context.Background(), {args}) }},",
                "\t\t},",
            ]
        )

    lines.extend(
        [
            "\t}",
            "",
            "\tfor _, tc := range cases {",
            "\t\tt.Run(tc.name, func(t *testing.T) {",
            "\t\t\tservice := " + constructor_call + '(rawStubRequester{',
            "\t\t\t\tt:         t,",
            "\t\t\t\twantPath:  tc.wantPath,",
            "\t\t\t\twantQuery: query,",
            f'\t\t\t\tpayload:   `{PAYLOAD}`,',
            "\t\t\t})",
            "",
            "\t\t\tresponse, err := tc.call(service)",
            "\t\t\tif err != nil {",
            '\t\t\t\tt.Fatalf("method returned error: %v", err)',
            "\t\t\t}",
            "",
            f'\t\t\tif got, want := string(response), `{PAYLOAD}`; got != want {{',
            '\t\t\t\tt.Fatalf("unexpected response: got %q want %q", got, want)',
            "\t\t\t}",
            "\t\t})",
            "\t}",
            "}",
        ]
    )

    return "\n".join(lines) + "\n"


def main() -> None:
    spec = json.loads(SPEC_PATH.read_text())
    grouped: dict[str, list[dict]] = defaultdict(list)

    for path, methods in spec["paths"].items():
        get = methods.get("get")
        if get is None or path in SKIP_PATHS:
            continue

        package = get["tags"][0].lower()
        grouped[package].append(
            {
                "path": path,
                "method_name": exported(get["operationId"]),
                "params": path_params(path),
            }
        )

    for package, methods in grouped.items():
        methods.sort(key=lambda item: item["method_name"])
        package_dir = ROOT / package
        package_dir.mkdir(parents=True, exist_ok=True)

        if package in {"user", "faction"}:
            (package_dir / "generated_raw.go").write_text(
                package_service_code(package, methods, include_constructor=False)
            )
            (package_dir / "generated_raw_test.go").write_text(
                package_test_code(package, methods, use_existing_constructor=True)
            )
            continue

        (package_dir / "service.go").write_text(
            package_service_code(package, methods, include_constructor=True)
        )
        (package_dir / "service_test.go").write_text(
            package_test_code(package, methods, use_existing_constructor=False)
        )


if __name__ == "__main__":
    main()
