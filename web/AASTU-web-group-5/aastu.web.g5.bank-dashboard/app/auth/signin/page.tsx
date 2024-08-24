"use client";
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import { useState, useEffect } from 'react';
import { signIn } from 'next-auth/react';
import Image from 'next/image';
import mainIcon from "/public/assets/icons/logo-card.png";
import Background from '@/public/images/background.jpg'
interface LoginFormData {
	userName: string;
	password: string;
}

const Login = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<LoginFormData>();
  const router = useRouter();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [currentCharIndex, setCurrentCharIndex] = useState(0);
  const [currentParagraphCharIndex, setCurrentParagraphCharIndex] = useState(0);
  const [showDescription, setShowDescription] = useState(false);
  const [showParagraph, setShowParagraph] = useState(false);
  const welcomeMessage = "Welcome to Our Bank Dashboard";
  const paragraphText = "Manage your finances with ease and confidence.";

  useEffect(() => {
    const descriptionTimer = setTimeout(() => {
      setShowDescription(true);
    }, 500); // Delay of 0.5 seconds for the description

    return () => clearTimeout(descriptionTimer);
  }, []);

  useEffect(() => {
    if (showDescription && currentCharIndex < welcomeMessage.length) {
      const charTimer = setTimeout(() => {
        setCurrentCharIndex(currentCharIndex + 1);
      }, 100); // Delay of 0.1 seconds for each character

      return () => clearTimeout(charTimer);
    } else if (currentCharIndex === welcomeMessage.length) {
      // Start displaying the paragraph one character at a time
      setShowParagraph(true);
    }
  }, [showDescription, currentCharIndex]);

  useEffect(() => {
    if (showParagraph && currentParagraphCharIndex < paragraphText.length) {
      const charTimer = setTimeout(() => {
        setCurrentParagraphCharIndex(currentParagraphCharIndex + 1);
      }, 100); // Delay of 0.1 seconds for each character in the paragraph

      return () => clearTimeout(charTimer);
    }
  }, [showParagraph, currentParagraphCharIndex]);

  const onSubmit = async (data: LoginFormData) => {
    setErrorMessage(null);

		try {
			const result = await signIn("credentials", {
				redirect: false,
				userName: data.userName,
				password: data.password,
			});

      if (result?.error) {
        setErrorMessage(result.error);
      } else if (result?.ok) {
        router.push('/');
      } else {
        setErrorMessage('Unexpected response from server.');
      }
    } catch (error: any) {
      setErrorMessage('An error occurred during login.');
      console.error("Login error:", error);
    }
  };

  return (
    <div  className="flex flex-col items-center min-h-screen bg-gradient-to-br from-blue-900 via-gray-800 to-blue-950 p-6">
      <div className="flex flex-col md:flex-row  items-center w-full max-w-5xl bg-white/10 backdrop-blur-lg p-8 sm:p-12 md:p-16 lg:p-20 rounded-2xl shadow-lg">
        <div className="w-full md:w-1/2 md:pr-12 mb-8 md:mb-0 text-center md:text-left">
          <div className="mb-6">
            <Image src={mainIcon} alt="BankDash Logo" className="h-24 w-24 mx-auto md:mx-0" />
            <h1 className="text-2xl font-bold text-white mt-4">Bank Dashboard</h1>
          </div>
          <h1 className="text-4xl font-extrabold text-white">
            {welcomeMessage.slice(0, currentCharIndex)}
          </h1>
          {showParagraph && (
            <p className="text-lg text-gray-300 mt-3">
              {paragraphText.slice(0, currentParagraphCharIndex)}
            </p>
          )}
        </div>
        <div className="w-full md:w-1/2">
          <h2 className="text-3xl sm:text-4xl md:text-5xl font-bold mb-8 sm:mb-10 md:mb-12 text-white text-center md:text-left">
            Login to Your Account
          </h2>
          {errorMessage && (
            <p className="text-red-500 mb-4 text-center md:text-left">{errorMessage}</p>
          )}
          <form onSubmit={handleSubmit(onSubmit)} className="w-full flex flex-col gap-6">
            <div className="w-full">
              <input
                {...register('userName', { required: 'Username is required' })}
                placeholder="Username"
                className="input-field w-full bg-gray-800 border border-gray-600 rounded-full p-4 text-white focus:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              {errors.userName && <p className="text-red-500 mt-2">{errors.userName.message}</p>}
            </div>

            <div className="w-full">
              <input
                {...register('password', { required: 'Password is required' })}
                type="password"
                placeholder="Password"
                className="input-field w-full bg-gray-800 border border-gray-600 rounded-full p-4 text-white focus:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              {errors.password && <p className="text-red-500 mt-2">{errors.password.message}</p>}
            </div>

            <button
              type="submit"
              className="w-full py-3 bg-gradient-to-r from-blue-600 to-blue-500 text-white font-semibold rounded-full shadow-lg hover:shadow-xl transition-transform transform hover:scale-105"
            >
              Login
            </button>
          </form>
          <p className="mt-8 text-center text-gray-400">
            Don &apos; t have an account? <a href="/auth/signup" className="text-blue-400 hover:text-blue-600 underline">Sign up</a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default Login;