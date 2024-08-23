import 'package:flutter/material.dart';

import '../../../product/presentation/widgets/text_field.dart';

class SignupPage extends StatefulWidget {
  SignupPage({super.key});

  @override
  State<SignupPage> createState() => _SignupPageState();
}

class _SignupPageState extends State<SignupPage> {
  final TextEditingController emailController = TextEditingController();

  final TextEditingController nameController = TextEditingController();

  final TextEditingController passwordController = TextEditingController();

  final TextEditingController confirmPasswordController =
      TextEditingController();

  bool _isChecked = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: SafeArea(
          child: Column(
            children: [
              Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: 20,
                ),
                child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      IconButton(
                        onPressed: () {
                          Navigator.pushNamed(context, '/signin_page');
                        },
                        icon: const Icon(
                          Icons.arrow_back_ios,
                          color: Color.fromARGB(255, 32, 68, 228),
                        ),
                      ),
                      Container(
                        padding: const EdgeInsets.symmetric(
                          horizontal: 10,
                        ),
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(10),
                          color: Colors.white,
                          border: Border.all(
                            color: const Color.fromARGB(255, 12, 62, 243),
                          ),
                          boxShadow: const [
                            BoxShadow(
                                color: Color.fromARGB(255, 195, 195, 195),
                                blurRadius: 10,
                                offset: Offset(0, 7))
                          ],
                        ),
                        child: const Text(
                          'ECOM',
                          style: TextStyle(
                            fontSize: 20,
                            fontWeight: FontWeight.bold,
                            color: Color.fromARGB(255, 33, 68, 243),
                            fontFamily: 'Caveat Brush',
                          ),
                        ),
                      ),
                    ]),
              ),
              const SizedBox(
                height: 40,
              ),
              const Text(
                'Create your account',
                style: TextStyle(
                  fontSize: 26,
                  fontWeight: FontWeight.w600,
                ),
              ),
              Container(
                  margin: const EdgeInsets.symmetric(horizontal: 42, vertical: 40),
                  child: Column(
                    children: [
                      MyTextField(
                        lable: 'Name',
                        lines: 1,
                        controller: nameController,
                        hint: 'ex: jon smith',
                      ),
                      MyTextField(
                        lable: 'Email',
                        lines: 1,
                        controller: emailController,
                        hint: 'ex: jon.smith@email.com',
                      ),
                      MyTextField(
                          lable: 'Password',
                          lines: 1,
                          controller: passwordController,
                          hint: '********',
                          obscureText: true),
                      MyTextField(
                          lable: 'Confirm Password',
                          lines: 1,
                          controller: confirmPasswordController,
                          hint: '********',
                          obscureText: true),
                      Row(
                        children: [
                          Checkbox(
                            value: _isChecked,
                            onChanged: (bool? value) {
                              setState(() {
                                _isChecked = value!;
                              });
                            },
                          ),
                          const Text(
                            'I understood the',
                            style: TextStyle(
                              fontSize: 15,
                              color: Color.fromARGB(255, 149, 148, 148),
                            ),
                          ),
                          GestureDetector(
                            onTap: () => {},
                            child: const Text('terms & policy.',
                                style: TextStyle(
                                  color: Color.fromARGB(255, 38, 80, 232),
                                )),
                          ),
                        ],
                      ),
                    ],
                  )),
              ElevatedButton(
                onPressed: () {},
                style: ButtonStyle(
                  backgroundColor: MaterialStateProperty.all(
                    const Color.fromARGB(255, 38, 80, 232),
                  ),
                  minimumSize: MaterialStateProperty.all(const Size(300, 50)),
                  shape: MaterialStateProperty.all(
                    RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                ),
                child: const Text(
                  'SIGN UP',
                  style: TextStyle(color: Colors.white, fontSize: 16),
                ),
              ),
              Container(
                margin: const EdgeInsets.only(left: 70, right: 60, top: 50),
                child: Row(
                  children: [
                    const Text(
                      ' Have an account? ',
                      style: TextStyle(
                          fontSize: 15,
                          color: Color.fromARGB(255, 149, 148, 148)),
                      textAlign: TextAlign.center,
                    ),
                    GestureDetector(
                      onTap: () {
                        Navigator.pushNamed(context, '/signin_page');
                      },
                      child: const Text('SIGN IN',
                          style: TextStyle(
                            color: Color.fromARGB(255, 38, 80, 232),
                          )),
                    )
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
