import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_8/features/product/presentation/bloc/product_event.dart';
import 'package:product_8/features/product/presentation/bloc/product_state.dart';
import 'package:product_8/features/product/presentation/pages/product_details_page.dart';


class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

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
        home: Detailspage(product: testProduct),
      ),
    );
  }

  testWidgets('displays product details', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.pumpAndSettle(); // Wait for UI rendering

    // Assert
    expect(find.text('Test Product'), findsOneWidget);
    expect(find.text('Test Description'), findsOneWidget);
    expect(find.text("\$100.0"), findsOneWidget);
  });

  testWidgets('DELETE button triggers DeleteProductEvent', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.pumpAndSettle();
    await tester.tap(find.text('DELETE'));
    await tester.pump(); // Trigger the event

    // Assert
    verify(() => mockProductBloc.add(DeleteProductEvent(id: testProduct.id))).called(1);
  });

  testWidgets('shows success Snackbar on successful deletion', (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable([ProductInitial(),ProductDeleteState()]),
      initialState: ProductInitial(),
    );

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.tap(find.text('DELETE'));
    await tester.pump(); // Trigger Snackbar
    await tester.pumpAndSettle(); // Wait for Snackbar to appear

    // Assert
    expect(find.text('deleted successfully'), findsOneWidget);
  });

  testWidgets('shows error Snackbar on deletion failure', (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable([ProductError()]),
      initialState: ProductInitial(),
    );

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.tap(find.text('DELETE'));
    await tester.pump(); // Trigger Snackbar
    await tester.pumpAndSettle(); // Wait for Snackbar to appear

    // Assert
    expect(find.text('Error deleting product'), findsOneWidget);
  });
}

// Fallback classes for Mocktail
class FakeProductEvent extends Fake implements ProductEvent {}
class FakeProductState extends Fake implements ProductState {}
