export default function Password({
  register,
  errors,
  value,
  placeHolder,
}: {
  register: any;
  errors: any;
  value: string;
  placeHolder?: string;
}) {
  const data = value.split(' ').join('').toLowerCase();
  return (
    <>
      <div className='my-3'>
        <label
          htmlFor={data}
          className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'
        >
          {value}
        </label>
        <input
          type='password'
          id={data}
          className='w-full font-[500] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
          {...register(data, {
            required: 'Password is required',
            minLength: {
              value: 8,
              message: 'Minimum length should be 8',
            },
            // pattern: {
            //   value: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/,
            //   message:
            //     'Password should contain at least one uppercase letter, one lowercase letter, and one number',
            // },
          })}
          placeholder={placeHolder}
        />
        {errors[data] && (
          <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
            {errors[data].message}
          </p>
        )}
      </div>
    </>
  );
}
