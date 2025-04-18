
from fastapi import APIRouter
from pydantic import BaseModel
from services.agents.anality_message_commit import CommitAnalyzer
from services.bot_models import botModel


class CommitMessageAnalizer(BaseModel):
    code_changes: str
    description: str
    tag: str
    language: str


router = APIRouter()


@router.post("/commit-analiszer/")
async def analizer_commit(data: CommitMessageAnalizer):
    gemini_model = botModel()
    print(f"Usando provider: {gemini_model.provider}, model: {gemini_model.model}")

    commitGeneration = CommitAnalyzer(llm_model=gemini_model)

    resp_ia = await commitGeneration.generate_commit_message(
        code_changes=data.code_changes,
        description=data.description,
        tag=data.tag,
        language=data.language,
    )

    return {"response": resp_ia}
