from fastapi import FastAPI
from routes.router import register_routers

app = FastAPI(title="Commit-IA", root_path="/commit-ia")


register_routers(app)
