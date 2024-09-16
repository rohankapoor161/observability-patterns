# Contributing to Observability Patterns

Thanks for considering contributing! These patterns come from production experience—real-world lessons from real incidents.

## What We're Looking For

- **Language implementations** — Go is primary, but Python/TypeScript examples welcome
- **Anti-patterns** — "We tried this and it failed" is as valuable as "this works"
- **Tool integrations** — OpenTelemetry, Prometheus, etc.
- **Case studies** — Real scenarios with anonymized details

## Format

Each pattern should include:
1. **Problem** — What pain does this solve?
2. **Solution** — Clear, practical implementation
3. **When to use** — Not everything belongs everywhere
4. **Code examples** — Working examples in `/examples/`

## Process

1. Open an issue to discuss the pattern
2. Fork and create a feature branch
3. Add your pattern with tests
4. Submit PR with clear description

## Code Style

- Go: Follow `gofmt` and standard Go conventions
- Include tests for all code examples
- Document exported functions

## Questions?

Open an issue or reach out: rohan.kapoor@example.com

---

*Remember: the goal is patterns that survive contact with production.*