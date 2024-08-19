import Link from 'next/link';

export default function Home() {
  return (
    <div className='w-screen min-h-screen flex justify-center items-center bg-slate-100'>
      <div>
        <div className='my-5'>
          <Link href='/api/auth/signin'>
            <button className='min-w-24 px-4 py-2 bg-indigo-900 text-xl text-gray-50 rounded-2xl'>
              SignIn
            </button>
          </Link>
        </div>
        <div>
          <Link href='/api/auth/signup'>
            <button className=' min-w-24 px-4 py-2 bg-indigo-900 text-xl text-gray-50 rounded-2xl'>
              SignUp
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
}
