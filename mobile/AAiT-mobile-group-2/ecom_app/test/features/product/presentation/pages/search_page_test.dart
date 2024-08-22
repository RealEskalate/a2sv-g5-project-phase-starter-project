import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecom_app/features/product/presentation/pages/search_page.dart';
import 'package:ecom_app/features/product/presentation/widgets/product_card.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
    HttpOverrides.global = null;
  });

  Widget _makeTestableWidget(Widget body) {
    return BlocProvider<ProductBloc>.value(
      value: mockProductBloc,
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
    await widgetTester.pumpWidget(_makeTestableWidget(const SearchPage()));

    expect(find.text('Search Product'), findsOneWidget);

    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });
  testWidgets('SearchPage should have ProductCard',
      (WidgetTester tester) async {
    //arrange
    when(() => mockProductBloc.state)
        .thenReturn(LoadAllProductState(products: testProductEntityList));

    //act
    await tester.pumpWidget(_makeTestableWidget(const SearchPage()));

    expect(find.byType(ProductCard), findsWidgets);
  });
  testWidgets('SearchPage shows error message when state is error',
      (WidgetTester tester) async {
    //arrange
    when(() => mockProductBloc.state)
        .thenReturn(ProductErrorState(message: 'Test Error Message'));

    //act
    await tester.pumpWidget(_makeTestableWidget(const SearchPage()));
    await tester.pumpAndSettle();

    expect(find.text('Error: Test Error Message'), findsOneWidget);
  });
}
