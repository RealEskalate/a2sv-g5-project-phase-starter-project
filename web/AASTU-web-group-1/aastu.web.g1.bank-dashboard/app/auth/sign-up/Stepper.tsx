import React from "react";

const Stepper = ({ current }: { current: string }) => {
  const steps = [
    { label: "Step 1", value: "Personal Information" },
    { label: "Step 2", value: "Address & Profile Picture" },
    { label: "Step 3", value: "Preferences" },
  ];

  return (
    <ol className="flex items-center w-full text-sm text-gray-500 font-medium sm:text-base">
      {steps.map((step, index) => {
        const isActive = current === step.value;
        const isCompleted = index < steps.findIndex((s) => s.value === current);
        const circleClass = isActive
          ? "bg-[#343C6A] border-indigo-200 text-white"
          : isCompleted
          ? "bg-gray-100 border-gray-200 text-gray-600"
          : "bg-gray-100 border-gray-200 text-gray-600";
        const textClass = isActive ? "text-[#343C6A]" : "text-gray-600";
        const afterClass =
          index < steps.length - 1
            ? "after:content-['/'] sm:after:hidden after:mx-2"
            : "";

        return (
          <li
            key={index}
            className={`flex md:w-full items-center ${textClass} ${
              index < steps.length - 1
                ? "sm:after:content-[''] after:w-full after:h-1 after:border-b after:border-gray-200 after:border-1 after:hidden sm:after:inline-block after:mx-4 xl:after:mx-8"
                : ""
            }`}
          >
            <div
              className={`flex items-center whitespace-nowrap ${afterClass}`}
            >
              <span
                className={`w-6 h-6 ${circleClass} border rounded-full flex justify-center items-center mr-3 text-sm lg:w-10 lg:h-10`}
              >
                {index + 1}
              </span>{" "}
              {step.label}
            </div>
          </li>
        );
      })}
    </ol>
  );
};

export default Stepper;
