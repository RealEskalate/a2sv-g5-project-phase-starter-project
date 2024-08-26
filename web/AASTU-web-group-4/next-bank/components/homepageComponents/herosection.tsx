import Link from "next/link";

const Hero = () => {
  return (
    <div className="relative h-screen bg-fixed bg-center bg-cover bg-[url('/Images/bank-background.jpg')]">
      <div className="absolute inset-0 bg-blue-800 bg-opacity-50 flex flex-col items-center justify-center text-white">
        <h1 className="text-5xl font-bold mb-4">Welcome to Our Platform</h1>
        <p className="text-lg mb-6">We offer amazing services to boost your business</p>
        <Link href="/signup">
          <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Get Started
          </button>
        </Link>
      </div>
    </div>
  );
};

export default Hero;
