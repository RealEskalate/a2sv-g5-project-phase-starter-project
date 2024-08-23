import 'pages.dart';
// import '../../../../config/route/route.dart' as route;

class SignUpPage extends StatefulWidget {
  const SignUpPage({super.key});

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  var checkBoxValue = false;

  late TextEditingController _emailController;
  late TextEditingController _nameController;
  late TextEditingController _passwordController;
  late TextEditingController _confirmPasswordController;

  @override
  void initState() {
    super.initState();
    _emailController = TextEditingController();
    _nameController = TextEditingController();
    _passwordController = TextEditingController();
    _confirmPasswordController = TextEditingController();
  }

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    _confirmPasswordController.dispose();
    _nameController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    const color = Color(0xFF3F51F3);
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is SignedUpState) {
          showCustomSnackBar(context, 'Successfully Signed UP');
          Navigator.of(context).pop();
        }
      },
      child: SafeArea(
        child: Scaffold(
          appBar: AppBar(
            automaticallyImplyLeading: false,
            leading: IconButton(
                onPressed: () {
                  Navigator.of(context).pop();
                },
                icon: const Icon(
                  Icons.arrow_back_ios_new,
                  color: color,
                  size: 30,
                )),
            actions: const [
              Row(
                // mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  EcomLogo(
                    fontSize: 24,
                    width: 60,
                    height: 40,
                  ),
                  SizedBox(
                    width: 20,
                  )
                ],
              )
            ],
          ),
          body: SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.all(20),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text(
                    'Create your account',
                    style: TextStyle(fontSize: 27),
                  ),
                  const SizedBox(
                    height: 40,
                  ),
                  CustomTextField(
                    controller: _nameController,
                    field: 'Name',
                    hintText: 'ex:jon.smith@email.com',
                  ),
                  CustomTextField(
                    controller: _emailController,
                    field: 'Email',
                    hintText: 'ex:jon.smith@email.com',
                  ),
                  CustomTextField(
                    controller: _passwordController,
                    obscureText: true,
                    field: 'Password',
                    hintText: '*********',
                  ),
                  CustomTextField(
                    controller: _confirmPasswordController,
                    obscureText: true,
                    field: 'Confirm password',
                    hintText: '*********',
                  ),
                  const SizedBox(
                    height: 30,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Transform.scale(
                        scale: 0.8,
                        child: Checkbox(
                          value: checkBoxValue,
                          onChanged: (value) {
                            if (value != null) {
                              setState(() {
                                checkBoxValue = value;
                              });
                            }
                          },
                        ),
                      ),
                      const CustomText(
                        text: 'I understood the ',
                        fontSize: 12,
                      ),
                      TextButton(
                          onPressed: () {},
                          child: const CustomText(
                            text: 'terms & policy.',
                            color: color,
                            fontSize: 12,
                          ))
                    ],
                  ),
                  BlocBuilder<AuthBloc, AuthState>(
                    builder: (context, state) {
                      final isLoading = state == LoadingState();
                      return CustomOutlinedButton(
                        onPressed: isLoading
                            ? null
                            : () {
                                if (checkBoxValue) {
                                  context.read<AuthBloc>().add(SignUpEvent(
                                      name: _nameController.text,
                                      email: _emailController.text,
                                      password: _passwordController.text,
                                      repeatedPassword:
                                          _confirmPasswordController.text));
                                } else {
                                  showCustomSnackBar(
                                      context, 'Agree with the terms & policy');
                                }
                              },
                        text:
                            state == LoadingState() ? 'Loading...' : 'SIGN UP',
                        backgroundColor: color,
                        color: Colors.white,
                      );
                    },
                  ),
                  const SizedBox(
                    height: 150,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const CustomText(
                        text: 'Have an account?',
                        color: Color(0xFF888888),
                        fontSize: 16,
                      ),
                      TextButton(
                          onPressed: () {
                            Navigator.of(context).pop();
                          },
                          child: const CustomText(
                            text: 'SIGN IN',
                            color: color,
                            fontSize: 16,
                          ))
                    ],
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
