from langchain.agents import  tool
from langchain.utilities import DuckDuckGoSearchAPIWrapper

@tool
def set_decision(grade, reason) -> None:
    '''
    this tool simple returns its own argument (decision)
    if the post is valid it will be called with argument of "True"
    otherwise it is called with argument of 'False'
    so call this function with "True" or "False"
    
    Args:
        grade : (int) (the grade of the given post from 0 to 100 )
        reason: (str) (detailed reason for the decision)
    returns:
        None
    
    '''
    from post_moderator import decision_state
    decision_state.message = reason
    decision_state.grade = grade
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
    it generates and returns image uri
    the image must be detailed and informative. if the prompt is not detailed the ai will generate random image which negatively affects your blog so be careful here
    args:
        prompt: (str) (what you want to generate) detailed description of the image to be generated
    returns:
        image url: (str) (the url of the image)
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
        return image_url
    
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