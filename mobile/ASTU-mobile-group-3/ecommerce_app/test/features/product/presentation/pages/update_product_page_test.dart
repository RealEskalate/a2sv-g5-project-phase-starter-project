import 'package:ecommerce_app/core/validator/validator.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/cubit/user_input_validation_cubit.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_states.dart';
import 'package:ecommerce_app/features/product/presentation/pages/update_product_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

import '../../../auth/presentation/page/init_files.dart';

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
  testWidgets('update product page ...', (tester) async {
    await tester.pumpWidget(pageTester(UpdateProductPage()));

    final name = find.byKey(const Key(InputDataValidator.name));

    final catagory = find.byKey(const Key(InputDataValidator.catagory));
    final price = find.byKey(const Key(InputDataValidator.price));
    final desc = find.byKey(const Key('Description'));
    final button = find.byType(FilledButton);
    expect(name, findsOneWidget);
    expect(catagory, findsNothing);
    expect(price, findsOneWidget);
    expect(desc, findsOneWidget);
    expect(button, findsOneWidget);
  });
}
