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
import 'package:product_8/features/product/presentation/pages/search_page.dart';
import 'package:product_8/features/product/presentation/widgets/product_card.dart';

// Mock classes
class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
    HttpOverrides.global = null;
    mockProductBloc = MockProductBloc();
  });

  Widget _makeTestableWidget() {
    return BlocProvider<ProductBloc>.value(
      value: mockProductBloc,
      child: const MaterialApp(home: SearchPage()),
    );
  }

  testWidgets('displays loading indicator when state is ProductLoading',
      (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenAnswer((_) => ProductLoading());

    await tester.pumpWidget(_makeTestableWidget());
    

    // Assert
    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });

  testWidgets('displays products when state is ProductLoaded',
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
    expect(find.byType(ProductCard), findsNWidgets(testProducts.length));
  });

  testWidgets('displays error message when state is ProductError',
      (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductError());

    // Act
    await tester.pumpWidget(_makeTestableWidget());
   

    // Assert
    expect(find.text('Error loading products'), findsOneWidget);
  });
  

}
