import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecom_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecom_app/features/product/presentation/pages/home_page.dart';
import 'package:ecom_app/features/product/presentation/widgets/product_card.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}

void main() {
  late MockProductBloc mockProductBloc;
  late MockAuthBloc mockAuthBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
    mockAuthBloc = MockAuthBloc();
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());
    when(() => mockAuthBloc.state).thenReturn(AuthInitial());
    HttpOverrides.global = null;
  });

  Widget _makeTestableWidget(Widget body) {
    return MultiBlocProvider(
      providers: [
        BlocProvider<ProductBloc>.value(
          value: mockProductBloc,
        ),
        BlocProvider<AuthBloc>.value(
          value: mockAuthBloc,
        ),
      ],
      child: MaterialApp(
        home: body,
      ),
    );
  }

  const testProductEntityList = [
    Product(
        id: '1',
        name: 'Test Pineapple',
        description: 'A yellow pineapple for the summer',
        imageUrl: 'pineapple.jpg',
        price: 5.33)
  ];

  testWidgets('state should have a loading circle', (widgetTester) async {
    //arrange
    when(() => mockProductBloc.state).thenAnswer((_) => ProductLoading());

    //act
    await widgetTester.pumpWidget(_makeTestableWidget(const HomePage()));

    expect(find.text('Aug 7, 2024'), findsOneWidget);

    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });
  testWidgets('HomePage should have ProductCard', (WidgetTester tester) async {
    //arrange
    when(() => mockProductBloc.state)
        .thenReturn(LoadAllProductState(products: testProductEntityList));

    //act
    await tester.pumpWidget(_makeTestableWidget(const HomePage()));

    expect(find.byType(ProductCard), findsWidgets);
  });
  testWidgets('Homepage shows error message when state is error',
      (WidgetTester tester) async {
    //arrange
    when(() => mockProductBloc.state)
        .thenReturn(ProductErrorState(message: 'Test Error Message'));

    //act
    await tester.pumpWidget(_makeTestableWidget(const HomePage()));
    await tester.pumpAndSettle();

    expect(find.text('Error: Test Error Message'), findsOneWidget);
  });
}
