async function LogIn({ userName, password }: any) {
    const data = { userName, password }; 
  
    try {
      const response = await fetch(
              
              "https://bank-dashboard-latest.onrender.com/auth/login",
        
              {
                method: "POST",
                headers: {
                  "Content-Type": "application/json",
                },
                body: JSON.stringify(data), 
                
              }
            );
      console.log(data);
  
      if (!response.ok) {
        throw new Error("Failed to log in");
      }
  
      const resData = await response.json(); 
      return resData.data;
    } catch (error) {
      console.error("Error in login.ts:", error);
      throw error;
    }
  }
  export default LogIn;