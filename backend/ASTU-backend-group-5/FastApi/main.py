from fastapi import FastAPI, HTTPException
from fastapi.responses import StreamingResponse

from pydantic import BaseModel
from post_moderator import moderator_agent, decision_state


from blog_writter import blog_assistant,blog

import os
from dotenv import load_dotenv
load_dotenv()

os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_KEY")



class BlogPost(BaseModel):
    title: str
    content: str

class Q(BaseModel):
    query: str
    chat_id: str


app = FastAPI()
@app.post("/validate_post/")
async def validate_post_endpoint(post: BlogPost):
    print(post)
    try:
        _ = moderator_agent.invoke("Title: " + post.title + " Content: " + post.content)
        return {"grade": decision_state.grade, "message": decision_state.message}
    except Exception as e:
        print(str(e))
        raise HTTPException(status_code=500, detail=str(e))
    
    
@app.post("/blog_assistant/")
async def stream_blog(query: Q):
    print(query)
    try:
        blog_assistant.run(query)
        blog.content = " ".join(blog.content)
        return blog
    except Exception as e:
        print(str(e))
        raise HTTPException(status_code=500, detail=str(e))




if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
