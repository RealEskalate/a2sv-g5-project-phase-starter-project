async function signInUser({ userName, password }: any) {
  const data = { userName, password }; // Your data object

  try {
    const response = await fetch("https://bank-dashboard-xx3n.onrender.com/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data), // Convert the data object to a JSON string
    });

    if (!response.ok) {
      // Check if the response is JSON and parse it accordingly
      const errorText = await response.text(); // Read response as plain text
      try {
        const errorJson = JSON.parse(errorText); // Try parsing the error text as JSON
        throw new Error(`Failed to sign in: ${errorJson.message || errorJson.error || "Unknown error"}`);
      } catch {
        // If parsing fails, throw the plain text error message
        throw new Error(`Failed to sign in: ${errorText}`);
      }
    }

    const resData = await response.json(); // Parse the JSON response if status is OK
    return resData.data;
  } catch (error) {
    console.error("Error in signInUser:", error);
    throw error; // Re-throw error for the caller to handle
  }
}

export default signInUser;
