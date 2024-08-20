"use client";
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { signIn } from 'next-auth/react';

interface LoginFormData {
  userName: string;
  password: string;
}

const Login = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<LoginFormData>();
  const router = useRouter();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const onSubmit = async (data: LoginFormData) => {
    setErrorMessage(null); // Reset error message

    try {
      const result = await signIn('credentials', {
        redirect: false,
        userName: data.userName,
        password: data.password,
      });

      if (result?.error) {
        setErrorMessage(result.error);
      } else if (result?.ok) {
        // Redirect to home page if sign-in was successful
        router.push('/');
      } else {
        // Handle unexpected cases
        setErrorMessage('Unexpected response from server.');
      }
    } catch (error: any) {
      setErrorMessage('An error occurred during login.');
      console.error("Login error:", error);
    }
  };

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100 p-4">
      <div className="w-full sm:max-w-lg md:max-w-[70%] lg:max-w-[60%] bg-white p-8 sm:p-12 md:p-16 lg:p-20 rounded-lg shadow-md mt-4">
        <h1 className="text-2xl sm:text-3xl md:text-4xl font-bold mb-8 sm:mb-10 md:mb-12 text-purple-700 text-center">Login to Your Account</h1>
        {errorMessage && <p className="text-red-500 mb-4 text-center">{errorMessage}</p>}
        <form onSubmit={handleSubmit(onSubmit)} className="w-full flex flex-col items-center gap-4 sm:gap-6 md:gap-8">
          <div className="mb-4 w-full flex flex-col items-center">
            <input
              {...register('userName', { required: 'Username is required' })}
              placeholder="Username"
              className="input-field w-full sm:max-w-lg md:max-w-[70%] lg:max-w-[60%] bg-blue-50 border border-blue-300 rounded-full p-2 focus:bg-blue-100"
            />
            {errors.userName && <p className="error-text text-red-500 mt-2">{errors.userName.message}</p>}
          </div>

          <div className="mb-4 w-full flex flex-col items-center">
            <input
              {...register('password', { required: 'Password is required' })}
              type="password"
              placeholder="Password"
              className="input-field w-full sm:max-w-lg md:max-w-[70%] lg:max-w-[60%] bg-blue-50 border border-blue-300 rounded-full p-2 focus:bg-blue-100"
            />
            {errors.password && <p className="error-text text-red-500 mt-2">{errors.password.message}</p>}
          </div>

          <button
            type="submit"
            className="w-full sm:max-w-lg md:max-w-[70%] lg:max-w-[60%] mx-auto py-3 bg-purple-700 text-white font-semibold rounded-full transition duration-300"
          >
            Login
          </button>
        </form>

        <p className="mt-8 text-center">
          Don't have an account? <a href="/auth/signup" className="text-blue-500">Sign up</a>
        </p>
      </div>
    </div>
  );
};

export default Login;
