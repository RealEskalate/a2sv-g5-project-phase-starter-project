"use client";
import React, { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { DevTool } from '@hookform/devtools';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import SignUpFormTypes from '@/app/types/SignUpForm_types';
import Link from 'next/link';
import axios, { isAxiosError } from 'axios';
import { useRouter } from 'next/navigation';

// Define validation schema using Yup
const schema = yup.object().shape({
    name: yup.string().required('Name is required'),
    email: yup.string().email('Please enter a valid email address').required('Email Address is required'),
    password: yup.string().required('Password is required'),
    confirmPassword: yup.string()
        .oneOf([yup.ref('password'), null], 'Passwords must match')
});

const SignUpForm = () => {
    const { register, control, handleSubmit, watch, formState: { errors }, trigger } = useForm<SignUpFormTypes>({
        resolver: yupResolver(schema)
    });
    
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const router = useRouter();

    const password = watch("password", "");

    useEffect(() => {
        trigger("confirmPassword");
    }, [password, trigger]);

    const onSubmit = async (data: SignUpFormTypes) => {
        console.log(data)
    };

    return (
        <div>
            <form action=""></form>
            
        </div>
    );
};

export default SignUpForm;
