
from fastapi import APIRouter
from schemas.schemas import CommitMessageAnalizer
from services.agents.anality_message_commit import CommitAnalyzer
from services.bot_models import botModel

router = APIRouter()


@router.post("/commit-analyzer/")
async def analyze_commit(data: CommitMessageAnalizer):
    ai_model = botModel()
    print(f"Using provider: {ai_model.provider}, model: {ai_model.model}")

    commit_analyzer = CommitAnalyzer(llm_model=ai_model)

    analysis_result = await commit_analyzer.generate_commit_message(
        code_changes=data.code_changes,
        description=data.description,
        tag=data.tag,
        language=data.language,
    )

    return {"response": analysis_result}
