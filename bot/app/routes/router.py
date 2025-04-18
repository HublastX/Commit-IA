from fastapi import FastAPI
from routes.anality_commit import anality_commit_routes


def register_routers(app: FastAPI):
    app.include_router(anality_commit_routes.router, tags=["Anality Commit"])



