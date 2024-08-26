// SecurityForm.tsx
import React from 'react';
import { useForm, Controller } from 'react-hook-form';
import Input from '@/components/ui/Input';
import Toggle from '@/components/ui/Toggle';
import { changePassword } from '@/services/authentication';

const SecurityForm = () => {
  const { control, register, handleSubmit } = useForm();

  const onSubmit = async (data: any) => {
    try {
      const response = await changePassword(data);
      console.log('Change Password Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <h3 className="font-semibold">Two-factor Authentication</h3>
      <Controller
        control={control}
        name="twoFactorAuth"
        render={({ field }) => (
          <Toggle label="Enable or disable two-factor authentication" {...field} />
        )}
      />

      <h3 className="font-semibold">Change Password</h3>
      <div className="w-full max-w-xs">
        <Input label="Current Password" type="password" placeholder="******" {...register('currentPassword')} />
      </div>
      <div className="w-full max-w-xs">
        <Input label="New Password" type="password" placeholder="******" {...register('newPassword')} />
      </div>
      <div className="flex justify-center md:pt-20">
        <button type="submit" className="w-full max-w-xs mx-auto bg-blue-800 text-white py-2 rounded-md">
          Save
        </button>
      </div>
    </form>
  );
};

export default SecurityForm;
