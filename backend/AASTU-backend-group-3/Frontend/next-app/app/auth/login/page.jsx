"use client";
import React, { useState } from "react";
import { FaGoogle } from "react-icons/fa";
import { useForm } from "react-hook-form";

const LoginPage = () => {
  const { register, handleSubmit, reset } = useForm();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const LoginWithEmail = async (data) => {
    setLoading(true);
    setError("");
    setSuccess("");

    try {
      const response = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      const result = await response.json();

      if (response.ok) {
        setSuccess("Login successful!");
        // Store the token and handle the session
        localStorage.setItem("token", result.token);
        console.log(result);
      } else {
        setError(result.message || "Login failed. Please try again.");
      }
    } catch (err) {
      setError("An error occurred. Please try again later.");
    } finally {
      setLoading(false);
      reset(); // Reset the form fields
    }
  };

  const handleGoogleLogin = () => {
    setLoading(true);
    setError("");
    setSuccess("");

    // Open a new window for Google OAuth login
    const googleLoginWindow = window.open(
      "http://localhost:8080/auth/login/google",
      "_blank",
      "width=500,height=600"
    );

    // Event listener to capture the token after Google login
    window.addEventListener("message", (event) => {
      if (event.origin !== "http://localhost:8080") return; // Ensure the origin matches your backend
      const { token } = event.data;

      if (token) {
        setSuccess("Google login successful!");
        localStorage.setItem("token", token); // Store the token in localStorage
        console.log("Google token received:", token);
        googleLoginWindow.close(); // Close the login popup
      } else {
        setError("Google login failed.");
      }
      setLoading(false);
    });
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 rounded-md">
      <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
        <h2 className="text-2xl font-bold text-gray-800 mb-6 text-center">
          Login
        </h2>

        {/* Display success or error messages */}
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        {success && (
          <p className="text-green-500 text-center mb-4">{success}</p>
        )}

        {/* Email and Password Login */}
        <form onSubmit={handleSubmit(LoginWithEmail)} className="space-y-4">
          <div>
            <label
              htmlFor="email"
              className="block text-sm font-medium text-gray-700"
            >
              Email
            </label>
            <input
              id="email"
              type="email"
              {...register("email", { required: true })}
              className="mt-1 p-2 w-full border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Enter your email"
              disabled={loading}
            />
          </div>

          <div>
            <label
              htmlFor="password"
              className="block text-sm font-medium text-gray-700"
            >
              Password
            </label>
            <input
              id="password"
              type="password"
              {...register("password", { required: true })}
              className="mt-1 p-2 w-full border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Enter your password"
              disabled={loading}
            />
          </div>

          <button
            type="submit"
            className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 transition-colors"
            disabled={loading}
          >
            {loading ? "Logging in..." : "Login with Email"}
          </button>
        </form>

        <div className="mt-6 text-center">
          <p className="text-gray-600">Or</p>
          <button
            onClick={handleGoogleLogin}
            className="w-full bg-red-500 text-white py-2 rounded-md flex items-center justify-center space-x-2 hover:bg-red-600 transition-colors mt-4"
            disabled={loading}
          >
            <FaGoogle className="text-white" />
            <span>{loading ? "Logging in..." : "Login with Google"}</span>
          </button>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
