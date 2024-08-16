
async function signInUser({ userName, password }: any) {
    const data = { userName, password }; // Your data object

    try {
        const response = await fetch("https://bank-dashboard-6acc.onrender.com/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data), // Convert the data object to a JSON string
        });
        
        if (!response.ok) {
            throw new Error("Failed to sign in");
        }
        
        const resData = await response.json(); // Parse the JSON response
        return resData.data;
    } catch (error) {
        console.error("Error in signInUser:", error);
        throw error;
    }
}
export default signInUser