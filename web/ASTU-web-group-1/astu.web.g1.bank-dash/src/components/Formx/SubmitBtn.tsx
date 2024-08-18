import Link from 'next/link';

export default function SubmitBtm({ value, account }: { value: string; account: boolean }) {
  return (
    <>
      <button
        type='submit'
        className='w-full bg-indigo-900 text-white rounded-3xl py-2 font-epilogue font-[700] mt-5 hover:bg-indigo-800 transition-all duration-500'
      >
        {value}
      </button>
      <p className='text-sm font-epilogue font-medium text-slate-400 mt-2 mx-2'>
        {account ? (
          <>
            {`Already have an account?`}
            <Link href='/login'>
              <span className='text-indigo-800 font-[700] ml-1'>Log In</span>
            </Link>
          </>
        ) : (
          <>
            {`Don't have an account?`}
            <Link href='/signup'>
              <span className='text-indigo-800 font-[700] ml-1'>Sign Up</span>
            </Link>
          </>
        )}
      </p>
    </>
  );
}
