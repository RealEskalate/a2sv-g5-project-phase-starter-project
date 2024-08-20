from langchain.agents import  tool
from langchain.utilities import DuckDuckGoSearchAPIWrapper

@tool
def set_decision(decision, reason):
    '''
    this tool simple returns its own argument (decision)
    if the post is valid it will be called with argument of "True"
    otherwise it is called with argument of 'False'
    so call this function with "True" or "False"
    
    Args:
        decision : (bool) (wheather the given post is valid or not)
        reason: (str) (detailed reason for the decision)
    returns:
        decision : (bool) the argument passed to the function
        reason : (str) (detailed reason for the decision)
    
    '''
    from post_moderator import decision_state
    decision_state.valid = decision
    decision_state.message = reason
    return decision


@tool
def duck_duck_go_search(query):
    '''
    useful for when you need to answer questions about current events, trends, or advancements.
    
    args:
        query: (str) (what you want to search for)
    returns:
        result: (str) (the result of the search)
    '''
    search = DuckDuckGoSearchAPIWrapper()
    result = search.run(query)
    
    return result
    
    
    
class Blog:
    def __init__(self, title: str, content: list):
        self.title = title
        self.content = content
blog = Blog("", [])


@tool
def generate_image(prompt: str) -> str:
    """Useful for when you need to generate an image. for your blog post
    it generates and sets the image to the blog posts content
    the image must be detailed and informative. if the prompt is not detailed the ai will generate random image which negatively affects your blog so be careful here
    args:
        prompt: (str) (what you want to generate) detailed description of the image to be generated
    returns:
        status: (str) (the status of the image generation) 
    """
    try:
        from openai import OpenAI
        client = OpenAI()

        response = client.images.generate(
        model="dall-e-3",
        prompt= prompt,
        size="1024x1024",
        quality="standard",
        n=1,
        )

        image_url = response.data[0].url
        blog.content.append({"type" : "image", "url" : image_url}) # adds image_url
        image_count = 0
        for content in blog.content:
            if content["type"] == "image":
                image_count += 1
                
        res =  "image generated and added to the blog content " + "now you have generated " + str(image_count) + " images to your blog post"
        if image_count > 4:
            res += "you cant generate image anymore"
        return res
    
    except:
        return "failed to generate image"


@tool
def set_title(title: str) -> str:
    """
    used to set the title of the blog post. probabliy only called once the chain starts
    
    args:
        title: (str) (the title of the blog post)
    returns:
        status: (str)
    """
    blog.title = title
    return "title set success fully now the title of the blog is " + title + "keep on generating your amazing blog"

@tool
def add_content(content: str) -> str:
    """
    adds content to the blog post
    the content can be a code, imae url or text
    the content must be detailed and informative
    use formatting to make it look nice
    args:
        content: (str) (the content of the blog post)
    returns:
        status: (str)
    """
    blog.content.append({"type" : "text", "content" : content})
    content_length = len(blog.content)
    char_count = 0
    for content in blog.content:
        if content["type"] == "text":
            char_count += len(content["content"])
    return "content added successfully " + "now you blog has " + str(char_count) + " characters and " + str(content_length) + " content blocks " + "keep on generating your amazing blog or if you feel you are generated enough stop"

# @tool
# def generate_blog_outline(blog_description: str) -> str:
#     """
#     generates an outline for the blog post
#     you always use this tool before you generate the blog post
#     args:
#         blog_description: (str) (the description of the blog post)
#     returns:
#         outline: (str) (the outline of the blog post)
#     """
#     return "outline generated successfully "


# @tool
# def add_code(code: str) -> str:
#     """
#     adds code to the blog post
#     args:
#         code: (str) (the code of the blog post)
#     """
#     blog.content.append({"type" : "code", "content" : code})
    
#     return "code added successfully"