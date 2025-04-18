from pydantic import BaseModel


class CommitMessageAnalizer(BaseModel):
    code_changes: str
    description: str
    tag: str
    language: str
