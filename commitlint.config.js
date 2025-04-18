module.exports = {
    extends: ["@commitlint/config-conventional"],
    rules: {
        "type-enum": [
            2,
            "always",
            [
                "feat",
                "fix",
                "chore",
                "docs",
                "style",
                "refactor",
                "test",
                "revert",
                "cleanup",
                "build",
                "remover",
            ],
        ],
        "type-case": [2, "always", "lowerCase"],
        "scope-case": [2, "always", "lowerCase"],
        "subject-empty": [2, "never"],
        "subject-full-stop": [2, "never", "."],
        "scope-empty": [2, "never"],
    },
};
