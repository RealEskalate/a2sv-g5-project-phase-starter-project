import '../../../../config/route/route.dart' as route;
import 'pages.dart';

class SignInPage extends StatefulWidget {
  const SignInPage({super.key});

  @override
  State<SignInPage> createState() => _SignInPageState();
}

class _SignInPageState extends State<SignInPage> {
  late TextEditingController _emailController;
  late TextEditingController _passwordController;

  @override
  void initState() {
    super.initState();
    _emailController = TextEditingController();
    _passwordController = TextEditingController();
  }

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    const color = Color(0xFF3F51F3);
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is SignedInState) {
          context.read<AuthBloc>().add(GetUserEvent());
        } else if (state is UserIsReady) {
          Navigator.pushNamedAndRemoveUntil(
              context, route.bottomNav, (Route<dynamic> route) => false);
        } else if (state is ErrorState) {
          showCustomSnackBar(context, state.message);
        }
      },
      child: Scaffold(
        body: SafeArea(
            child: Padding(
          padding: const EdgeInsets.all(20),
          child: SingleChildScrollView(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const EcomLogo(
                  width: 120,
                  fontSize: 48,
                ),
                const SizedBox(height: 80),
                const Text(
                  'Sign into your account',
                  style: TextStyle(fontSize: 27),
                ),
                const SizedBox(
                  height: 20,
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
                const SizedBox(
                  height: 30,
                ),
                BlocBuilder<AuthBloc, AuthState>(
                  builder: (context, state) {
                    return CustomOutlinedButton(
                      onPressed: () {
                        context.read<AuthBloc>().add(
                              SignInEvent(
                                email: _emailController.text,
                                password: _passwordController.text,
                              ),
                            );
                      },
                      text: state == LoadingState() ? 'Loading...' : 'SIGN IN',
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
                      text: 'Don’t have an account?',
                      color: Color(0xFF888888),
                      fontSize: 16,
                    ),
                    TextButton(
                        onPressed: () {
                          Navigator.of(context).push(
                            MaterialPageRoute(
                              builder: (context) => const SignUpPage(),
                            ),
                          );
                        },
                        child: const CustomText(
                          text: 'SIGN UP',
                          color: color,
                          fontSize: 16,
                        ))
                  ],
                ),
              ],
            ),
          ),
        )),
      ),
    );
  }
}
