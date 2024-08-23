import 'package:bloc_test/bloc_test.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecommerce_app/features/auth/presentation/page/login_page.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
// import 'package:mockito/mockito.dart';

class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}

class MockInputValidationCubit extends MockCubit<InputValidationState>
    implements InputValidationCubit {}

void main() {
  late MockAuthBloc mockAuthBloc;

  late MockInputValidationCubit mockInputValidationCubit;

  setUp(() {
    mockAuthBloc = MockAuthBloc();
    mockInputValidationCubit = MockInputValidationCubit();
  });

  Future<void> pumpableWidget(WidgetTester tester) async {
    await tester.pumpWidget(MaterialApp(
      home: MultiBlocProvider(providers: [
        BlocProvider(create: (_) => mockAuthBloc),
        BlocProvider(create: (_) => mockInputValidationCubit),
      ], child: LoginPage()),
    ));
  }

  testWidgets('the widgets must appear on the ', (tester) async {
    when(() => mockAuthBloc.state).thenReturn(AuthInitial());

    when(() => mockInputValidationCubit.state)
        .thenReturn(InputValidationInitial());

    await pumpableWidget(tester);
  });
}
