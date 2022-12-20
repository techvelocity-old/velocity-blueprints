from fastapi import FastAPI
from fastapi.responses import  HTMLResponse

app = FastAPI()

@app.get('/')
async def index():
    return HTMLResponse("<h1> Velocity Env Sync Example </h1>")


    