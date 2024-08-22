import { FieldErrors, UseFormRegister } from "react-hook-form";

interface contactForm {
  name: string;
  email: string;
  date: string;
  PermanentAddress: string;
  postal: string;
  username: string;
  password: string;
  CurrentAddress: string;
  city: string;
  country: string;
}

export default function Input({
  field,
  namee,
  placeholder,
  regex,
  regexMsg,
  requiredMsg,
  minLength,
  errors,
  register,
  mode,
}: {
  field: string;
  namee: keyof contactForm;
  placeholder: string;
  regex: RegExp;
  regexMsg: string;
  requiredMsg: string;
  minLength: number;
  errors: FieldErrors<contactForm>;
  register: UseFormRegister<contactForm>;
  mode: string;
}) {
  return (
    <div className="mb-4 flex flex-col gap-1">
      <label htmlFor={namee} className="text-[#232323]">
        {field}
      </label>

      <input
        className="shadow-sm rounded-md border hover:border-gray-300 focus:outline-none focus:shadow-md p-3 px-4 bg-inherit text-gray-500"
        type={field === "Password" ? "password" : "text"}
        placeholder={placeholder}
        id={namee}
        {...register(namee, {
          pattern: {
            value: regex,
            message: regexMsg,
          },
          required: requiredMsg,
          minLength: {
            value: minLength,
            message: `Minimum length is ${minLength} characters`,
          },
        })}
      />

      <p className="pl-2 text-red-500">{errors[namee]?.message}</p>
    </div>
  );
}
