import Contacts from "@/components/Welcome-landingPage/Contacts";
import Description from "@/components/Welcome-landingPage/Description";
import Navbar from "@/components/Welcome-landingPage/navbar";
import WelcomeDashboard from "@/components/Welcome-landingPage/WelcomeDashboard";

export default function Home() {
  const currentYear = new Date().getFullYear();
  return (
    <div className="w-full min-h-screen bg-white scroll-smooth">
      <Navbar />
      <div className="flex flex-col w-full pt-32 md:pt-32 items-center space-y-20 px-4 md:px-14 md:space-y-20 mb-20">
        <WelcomeDashboard />
        <Description />
      </div>
      <Contacts />
      <footer className="bg-slate-50 text-gray-800 text-center py-1 w-full">
        <p>&copy; {currentYear} Bank . All rights reserved.</p>
      </footer>
    </div>
  );
}
