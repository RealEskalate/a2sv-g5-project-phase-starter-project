import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
import 'package:mocktail/mocktail.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_8/features/product/presentation/bloc/product_event.dart';
import 'package:product_8/features/product/presentation/bloc/product_state.dart';
import 'package:product_8/features/product/presentation/pages/home_page.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
    HttpOverrides.global = null;
  });

  Widget _makeTestableWidget() {
    return BlocProvider<ProductBloc>.value(
      value: mockProductBloc,
      child: const MaterialApp(home: HomePage()),
    );
  }

  testWidgets(
      'text fields should trigger state to change from empty to loading',
      (WidgetTester widgetTester) async {
    // arrange
    when(() => mockProductBloc.state).thenAnswer((_) => ProductLoading());
    // act
    await widgetTester.pumpWidget(_makeTestableWidget());
    
    
    // assert
    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });

  testWidgets('HomePage shows products when ProductLoaded state is emitted',
      (WidgetTester tester) async {
    // Arrange
    final testProducts = [
      const Product(
          id: '1',
          name: 'Test Product',
          price: 20.0,
          description: 'jjk',
          imageUrl: 'https://www.google.com')
    ];
    when(() => mockProductBloc.state)
        .thenReturn(ProductLoaded(products: testProducts));

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.pumpAndSettle();

    // Assert
    expect(find.text('Available Products'), findsOneWidget);
    expect(find.text('Test Product'), findsOneWidget);
  });

  testWidgets('HomePage shows error message when ProductError state is emitted', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductError());

    // Act
    await tester.pumpWidget(_makeTestableWidget());
    await tester.pumpAndSettle();

    // Assert
    expect(find.text('Error loading products'), findsOneWidget);
  });


}
