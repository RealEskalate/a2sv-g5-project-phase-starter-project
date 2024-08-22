import { useState } from "react";

interface Step1Props {
  data: { [key: string]: any };
  onNext: (data: { [key: string]: any }) => void;
}

const Step1: React.FC<Step1Props> = ({ data, onNext }) => {
  const [formData, setFormData] = useState<{ [key: string]: any }>(data);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  return (
    <div>
      <h2>Step 1</h2>
      <input
        type="text"
        name="field1"
        value={formData.field1 || ""}
        onChange={handleChange}
        className="border p-2"
        placeholder="Field 1"
      />
      <button
        onClick={() => onNext(formData)}
        className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
      >
        Save and Next
      </button>
    </div>
  );
};

export default Step1;
