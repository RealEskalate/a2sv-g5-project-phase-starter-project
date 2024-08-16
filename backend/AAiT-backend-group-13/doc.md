ENDPOINTS:
    Signup:
        request data:
            First name
            last name
            user name
            email
            password

        response data:
            message

    Login
        request data:
            user name
            password

        response data
            access token 
            refresh token
            user name
            user id
            isAdmin
            first name
            last name
    
    forgot password:
        request data:
            email
        
        response message:
            messagae 

    reset password:
        request data:
            token payload in link
            password
        response:
            messagae

    logout:
        request data:
            ---
        response data:
            message

    promotion:
        request data:
            user name
        response data:
            message
    
    demotion:
        request data:
            user name 
        response data:
            message

    create blog:
        request data:
            title 
            content
            tag
        response data:
            blog

    get all blogs:
        request data:
            payload - cursor head
            sorting parameter
            filtering 
        response data:
            blogs

    blog update
        request data:
            updated details
        response data:
            message

    blog delete:
        request data:
            blog id
        response data:
            message

    popularity tracking:
        react:
            request data:
                blog id
            response data:
                ----


    edit profile:
        request data:
            updated info
        response data:
            messsage

    comment:
        request data:
            blog id
            content
        response:
            message

    get comments:
        request data:
            blog id
        response data:
            comments

    AI Integration:
        // PASS
    Third party Auth:
        //PASS



    

    


    

