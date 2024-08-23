import 'dart:developer';

import '../../../../config/route/route.dart' as route;
import '../../../product/presentation/bloc/product_bloc.dart';
import 'pages.dart';

class SplashPage extends StatefulWidget {
  const SplashPage({super.key});

  @override
  State<SplashPage> createState() => _SplashPageState();
}

class _SplashPageState extends State<SplashPage> {
  @override
  void initState() {
    super.initState();

    Timer(
      const Duration(seconds: 3),
      () {
        context.read<AuthBloc>().add(CheckSignedInEvent());
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    final height = MediaQuery.sizeOf(context).height;
    final width = MediaQuery.sizeOf(context).width;

    return MultiBlocListener(
      listeners: [
        BlocListener<AuthBloc, AuthState>(
          listener: (context, state) {
            if (state is SignedInState) {
              context.read<AuthBloc>().add(GetUserEvent());
            } else if (state is UserIsReady) {
              Navigator.pushReplacementNamed(context, route.homePage);
            } else if (state is SignInState) {
              Navigator.pushReplacementNamed(context, route.signInPage);
            }
          },
        ),
      ],
      child: Scaffold(
        body: Stack(
          children: [
            Positioned.fill(
              child: Image.asset(
                'lib/assets/images/splash_bg.jpg',
                fit: BoxFit.cover,
              ),
            ),
            Positioned.fill(
              child: Container(
                height: height,
                width: width,
                decoration: BoxDecoration(
                  gradient: LinearGradient(colors: [
                    const Color(0xFF3F51F3),
                    const Color(0xFF3F51F3).withOpacity(0.2)
                  ], begin: Alignment.bottomCenter, end: Alignment.topCenter),
                ),
                child: const Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    AnimatedEcomLogo(),
                    CustomText(
                      text: 'Ecommerce APP',
                      fontSize: 35,
                      color: Colors.white,
                    )
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
