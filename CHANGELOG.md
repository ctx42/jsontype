## v0.9.0 (Fri, 02 Oct 2026 20:49:21 UTC)
- test: isolate package-level registry in Test_Register.
- fix: unmarshal integers beyond 2^53 without precision loss.
- fix: unmarshal time.Duration written as nanoseconds.
- fix!: marshal Value stored in non-addressable places.
- fix: leave Value unchanged when envelope conversion fails.
- fix: ignore whitespace and parse bare numbers strictly in Unmarshal.
- fix: return errors instead of panicking on nil registry or Value.
- feat: accept options in FromMap and AsValue.
- fix: wrap ErrInvFormat in the invalid JSON token error.
- fix: prefix MarshalJSON errors with jsontype.
- refactor: initialize package-level registry in its declaration.
- fix: make the zero value Registry usable.
- refactor: move Unmarshal into jsontype.go.
- test: order UnmarshalJSON table test after its subject's test.
- test: prepare call arguments in the Given section.
- test: name expected error messages want.
- test: inline the expected map in Test_Value_Map.
- test: rename duplicate Registry.Register subtest.
- docs: fix grammar in doc comments.
- refactor: drop duplicate nil check from the package-level Register.
- refactor: drop unused converter binding in NewValue.
- docs: explain when the byte and rune type names occur.
- fix: unmarshal float32 values written by MarshalJSON.
- test: round-trip every built-in type through JSON.
- chore: bump convert to v0.11.0 and testing to v0.56.0.
- fix: accept byte, rune and custom type names in FromMap.
- test: assert the full invalid JSON error in Unmarshal.
- fix: use the dynamic type name in New for interface types.
- fix: reject a nil *Value and accept a Value in AsValue.
- fix: return ErrNilRegistry from NewValue also for a nil value.
- docs: document nil and numeric handling in Registry.Register.
- test: round-trip the byte and rune type names through JSON.
- docs: explain the role of the package-level registry in Register.
- docs: add examples for Unmarshal, NewValue and AsValue.
- test: inline the expected error in Test_numberConverter.
- test: inline the expected error in Test_Value_MarshalJSON.
- test: declare Test_Register arguments in call order.
- test: separate subjects with blank lines in jsontype_test.go.
- test: separate subjects with blank lines in registry_test.go.
- test: separate subjects with blank lines in converters_test.go.
- test: separate subjects with blank lines in helpers_test.go.
- test: separate subjects with blank lines in options_test.go.
- test: separate subjects with blank lines in all_test.go.
- style: separate multi-line switch cases with blank lines.
- docs: say what Option and Options configure.
- test: add When and Then markers to Test_registry.
- refactor: drop redundant empty map check in keyValue.
- test: separate subtests in Test_unmarshalEnvelope with a blank line.
- test: rename the internal package subtest of Test_New.
- test: drop UnmarshalJSON subtests duplicated by the tabular test.
- docs: note the lasting registration in ExampleRegister_custom.
- docs: cover the full API and refresh examples in README.

## v0.8.0 (Sun, 24 May 2026 19:16:31 UTC)
- feat!: emit transparent types as bare JSON values.
- perf: eliminate json.Unmarshal/Marshal on bool and float64 paths.
- fix: validate full token in Unmarshal bool and null fast paths.
- refactor: simplify token validation with string comparison.
- test: declare byte payloads in Given, pin error messages.
- refactor: standardize error message prefix to `jsontype:`.
- test(helpers): improve Unmarshal coverage to 100%.
- docs: polish README and doc comments.
- fix(FromMap): avoid double jsontype: prefix on NewValue errors.

## v0.7.0 (Fri, 01 May 2026 20:09:25 UTC)
- chore: Update to Go 1.26 and update dependencies.

## v0.6.2 (Sun, 26 Apr 2026 13:39:38 UTC)
- chore: update dependencies.

## v0.6.1 (Mon, 30 Mar 2026 08:48:11 UTC)
- feat!: Rename `jsontype.Unmarshal` to `Unmarshal` to reduce stutter.
- fix(jsontype): correct duplicate registration and missing docs.
- chore(jsontype): Strip email addresses from SPDX copyright headers across all files.
- chore: exclude CLAUDE.md from version control.

## v0.6.0 (Thu, 12 Feb 2026 14:00:16 UTC)
- feat: Add DefaultRegistry function returning default registry configuration.
- feat: Add UnmarshalJSON function which allows you to provide custom Registry.

## v0.5.1 (Tue, 03 Feb 2026 09:12:45 UTC)
- chore: Update dependencies.

## v0.5.0 (Thu, 29 Jan 2026 11:24:15 UTC)
- chore: Update dependencies, improve error messages, add examples, and enhance documentation.

## v0.4.1 (Tue, 27 Jan 2026 16:10:49 UTC)
- chore: Update dependencies and tests.

## v0.4.0 (Sat, 20 Dec 2025 22:06:30 UTC)
- Use `github.com/ctx42/convert` for type conversions.

## v0.3.0 (Fri, 19 Dec 2025 12:11:27 UTC)
- Use `github.com/ctx42/convert` for type conversions.

## v0.2.0 (Sat, 29 Nov 2025 21:15:08 UTC)
- Improve ease of use by adding new constructor functions.

## v0.1.0 (Fri, 28 Nov 2025 20:33:22 UTC)
- Initial commit.
- Initial implementation.
- Add documentation and examples.
- Add package comment.
- Update package comment and documentation.

