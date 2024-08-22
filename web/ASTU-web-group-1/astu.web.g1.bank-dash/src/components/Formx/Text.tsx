export default function Text({
  register,
  errors,
  value,
  placeHolder,
}: {
  register: any;
  errors: any;
  value: string;
  placeHolder: string;
}) {
  const data = value.split(' ').join('').toLowerCase();
  return (
    <div className='my-3'>
      <label htmlFor={data} className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'>
        {value}
      </label>
      <input
        className='w-full font-[600] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
        type='text'
        id={data}
        {...register(data, {
          required: `${value} is required`,
          minLength: {
            value: 5,
            message: `Minimum length should be 5`,
          },
        })}
        placeholder={placeHolder}
      />
      {errors[data] && (
        <p className='text-red-500 text-xs mt-1 font-poppins font-[550]'>{errors[data].message}</p>
      )}
    </div>
  );
}