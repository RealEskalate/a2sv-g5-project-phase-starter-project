<<<<<<< HEAD
import Link from "next/link";
=======
import Contacts from "@/components/Welcome-landingPage/Contacts";
import Description from "@/components/Welcome-landingPage/Description";
import Navbar from "@/components/Welcome-landingPage/navbar";
import WelcomeDashboard from "@/components/Welcome-landingPage/WelcomeDashboard";
>>>>>>> 5b4019f66bbe3b9bb435b8528477c191f8c84beb

export default function Home() {
  const currentYear = new Date().getFullYear();
  return (
<<<<<<< HEAD
    <div className="w-screen min-h-screen flex justify-center items-center bg-slate-100">
      <div>
        <div className="my-5">
          <Link href="/api/auth/signin">
            <button className="min-w-24 px-4 py-2 bg-indigo-900 text-xl text-gray-50 rounded-2xl">
              SignIn
            </button>
          </Link>
        </div>
        <div>
          <Link href="/api/auth/signup">
            <button className=" min-w-24 px-4 py-2 bg-indigo-900 text-xl text-gray-50 rounded-2xl">
              SignUp
            </button>
          </Link>
        </div>
=======
    <div className="w-full min-h-screen bg-slate-50 scroll-smooth">
      <Navbar />
      <div className="flex flex-col w-full pt-32 md:pt-32 items-center space-y-20 px-4 md:px-14 md:space-y-20 mb-20">
        <WelcomeDashboard />
        <Description />
>>>>>>> 5b4019f66bbe3b9bb435b8528477c191f8c84beb
      </div>
      <Contacts />
      <footer className="bg-gray-800 text-white text-center py-1 w-full">
        <p>&copy; {currentYear} Bank . All rights reserved.</p>
      </footer>
    </div>
  );
}
