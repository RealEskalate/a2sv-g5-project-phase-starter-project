import Contacts from '@/components/Welcome-landingPage/Contacts';
import WelcomeDashboard from '@/components/Welcome-landingPage/WelcomeDashboard';

export default function Home() {
  const currentYear = new Date().getFullYear();
  return (
    <div className='w-full min-h-screen bg-slate-50 scroll-smooth'>
      <WelcomeDashboard />
      <Contacts />
      <footer className='text-slate-800 mt-10 text-center py-1 w-full'>
        <p>&copy; {currentYear} Bank . All rights reserved.</p>
      </footer>
    </div>
  );
}
