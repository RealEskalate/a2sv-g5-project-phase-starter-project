from fastapi import FastAPI, HTTPException
from fastapi.responses import StreamingResponse

from pydantic import BaseModel
from post_moderator import moderator_agent, decision_state

from blog_assistance_enhanced import BlogAssistant
from tools import *




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
        is_valid = moderator_agent.invoke("Title: " + post.title + " Content: " + post.content)
        return {"is_valid": decision_state.valid, "message": decision_state.message}
    except Exception as e:
        print(str(e))
        raise HTTPException(status_code=500, detail=str(e))
    
    
@app.post("/blog_assistant/")
async def stream_blog(query: Q):
    print(query)
    assistant = BlogAssistant(tools=[duck_duck_go_search, generate_image, add_content, set_title])  
    try:
        assistant.run(query)
        return blog
    except Exception as e:
        print(str(e))
        raise HTTPException(status_code=500, detail=str(e))





if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
