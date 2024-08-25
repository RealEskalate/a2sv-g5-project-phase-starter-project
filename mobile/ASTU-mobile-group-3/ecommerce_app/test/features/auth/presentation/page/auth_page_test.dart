import 'package:ecommerce_app/core/validator/validator.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/cubit/user_input_validation_cubit.dart';
import 'package:ecommerce_app/features/auth/presentation/page/login_page.dart';
import 'package:ecommerce_app/features/auth/presentation/page/signup_page.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_states.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

import 'init_files.dart';

void main() {
  late MockAuthBloc mockAuthBloc;
  late MockUserInput mockUserInput;
  late MockInputValidator mockInputValidator;
  late MockProductBloc mockProductBloc;

  setUp(() {
    mockUserInput = MockUserInput();
    mockAuthBloc = MockAuthBloc();
    mockProductBloc = MockProductBloc();
    mockInputValidator = MockInputValidator();
  });

  Widget pageTester(Widget insert) {
    return MaterialApp(
      home: MultiBlocProvider(providers: [
        BlocProvider<AuthBloc>.value(
          value: mockAuthBloc,
        ),
        BlocProvider<UserInputValidationCubit>.value(
          value: mockUserInput,
        ),
        BlocProvider<InputValidationCubit>.value(
          value: mockInputValidator,
        ),
        BlocProvider<ProductBloc>.value(
          value: mockProductBloc,
        ),
      ], child: insert),
    );
  }

  setUp(() {
    when(() => mockAuthBloc.state).thenReturn(AuthInitial());
    when(() => mockUserInput.state).thenReturn(UserInputValidationInitial());
    when(() => mockInputValidator.state).thenReturn(InputValidationInitial());
    when(() => mockProductBloc.state).thenReturn(InitialState());
  });
  group('Log in page testing', () {
    testWidgets('Should find the button in log in page', (tester) async {
      await tester.pumpWidget(pageTester(LoginPage()));

      final button = find.byKey(const Key('SIGN IN'));
      final emailInput = find.byKey(const Key(InputDataValidator.email));
      final passwordInput = find.byKey(const Key(InputDataValidator.password));
      expect(emailInput, findsOneWidget);
      expect(passwordInput, findsOneWidget);
      expect(button, findsOneWidget);
    });
  });

  group('Sign up page', () {
    testWidgets('Should find necessary widgets on sign in page',
        (tester) async {
      await tester.pumpWidget(pageTester(SignUpPage()));

      final name = find.byKey(const Key(InputDataValidator.name));
      final email = find.byKey(const Key(InputDataValidator.email));
      final password = find.byKey(const Key(InputDataValidator.password));
      final confPass = find.byKey(const Key(InputDataValidator.confirmPass));

      final checkBox = find.byKey(const Key(InputDataValidator.checkBox));

      expect(name, findsOneWidget);
      expect(email, findsOneWidget);
      expect(password, findsOneWidget);
      expect(confPass, findsOneWidget);
      expect(checkBox, findsOneWidget);
    });
  });
}
