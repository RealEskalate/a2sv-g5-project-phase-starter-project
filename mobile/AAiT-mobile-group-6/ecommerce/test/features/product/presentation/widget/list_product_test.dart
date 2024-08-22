// ignore_for_file: prefer_const_constructors

import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_state.dart';
import 'package:ecommerce/features/product/presentation/screens/homepage.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
  });

  testWidgets('displays product list correctly', (WidgetTester tester) async {
    // Arrange
    when(mockProductBloc.state).thenReturn(AllProductsLoaded(
      products: const [
        Product(
          id: '1',
          name: 'Product 1',
          price: 10.0,
          description: 'Description 1',
          imageUrl: 'https://example.com/image1.png',
        ),
        Product(
          id: '2',
          name: 'Product 2',
          price: 20.0,
          description: 'Description 2',
          imageUrl: 'https://example.com/image2.png',
        ),
      ],
    ));

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: BlocProvider<ProductBloc>(
          create: (_) => mockProductBloc,
          child: const HomePage(),
        ),
      ),
    );

    await tester.pumpAndSettle(); // Wait for the widget tree to settle.

    // Assert
    expect(find.text('Product 1'), findsOneWidget);
    expect(find.text('Product 2'), findsOneWidget);
    expect(find.text('\$10.0'), findsOneWidget);
    expect(find.text('\$20.0'), findsOneWidget);
  });
}
