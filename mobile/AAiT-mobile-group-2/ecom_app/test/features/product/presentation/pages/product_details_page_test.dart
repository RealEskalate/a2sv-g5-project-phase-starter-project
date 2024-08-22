import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecom_app/features/product/presentation/pages/product_details_page.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

// Fallback classes for Mocktail
class FakeProductEvent extends Fake implements ProductEvent {}

class FakeProductState extends Fake implements ProductState {}

void main() {
  late MockProductBloc mockProductBloc;
  late Product testProduct;

  setUpAll(() {
    registerFallbackValue(FakeProductEvent());
    registerFallbackValue(FakeProductState());
  });

  setUp(() {
    mockProductBloc = MockProductBloc();
    HttpOverrides.global = null;
    testProduct = const Product(
      id: '1',
      name: 'Test Product',
      price: 100.0,
      description: 'Test Description',
      imageUrl: 'https://example.com/image.jpg',
    );
  });

  Widget _makeTestableWidget() {
    return BlocProvider<ProductBloc>.value(
      value: mockProductBloc,
      child: MaterialApp(
        routes: {
          '/home': (context) => const Scaffold(
              body: Text('Navigation Successful'),
            ),
        },
        home: ProductDetailsPage(product: testProduct),
      ),
    );
  }

  testWidgets('displays product details', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.pumpAndSettle();

    // Assert
    expect(find.text('Test Product'), findsOneWidget);
    expect(find.text('Test Description'), findsOneWidget);
    expect(find.text('\$100.0'), findsOneWidget);
  });

  testWidgets('DELETE button triggers DeleteProductEvent',
      (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.pumpAndSettle();
    await tester.tap(find.text('DELETE'));
    await tester.pump();

    // Assert
    verify(() => mockProductBloc.add(DeleteProductEvent(id: testProduct.id)))
        .called(1);
  });

  testWidgets('shows success Snackbar on successful deletion',
      (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable(
          [ProductInitialState(), ProductDeletedState()]),
      initialState: ProductInitialState(),
    );

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.tap(find.text('DELETE'));
    await tester.pump(); // Trigger Snackbar
    await tester.pumpAndSettle(); // Wait for Snackbar to appear

    // Assert
    expect(find.text('Product Deleted Successfully'), findsOneWidget);
  });

  testWidgets('shows error Snackbar on deletion failure',
      (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable([ProductErrorState(message: '')]),
      initialState: ProductInitialState(),
    );

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.tap(find.text('DELETE'));
    await tester.pump();
    await tester.pumpAndSettle();

    // Assert
    expect(find.text('Failed to delete the product'), findsOneWidget);
  });
}
