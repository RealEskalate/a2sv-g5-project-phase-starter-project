const Error = ({ message }: { message: string }) => {
    return (
      <div className="flex justify-center items-center h-screen">
        <div className="bg-red-100 text-red-600 border border-red-300 rounded-lg p-4">
          <div className="flex items-center  w-72 h-20">
            <svg
              className="w-15 h-10 mr-2"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M12 8v4m0 4h.01M21 12c0 5.523-4.477 10-10 10S1 17.523 1 12 5.477 2 11 2s10 4.477 10 10z"
              />
            </svg>
            <span className="font-heading">{message}</span>
          </div>
        </div>
      </div>
    );
  };
  
  export default Error;
  