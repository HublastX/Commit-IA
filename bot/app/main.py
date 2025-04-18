from fastapi import FastAPI
from routes.router import register_routers

app = FastAPI(title="Eduga API", root_path="/anality-pdf")


register_routers(app)
