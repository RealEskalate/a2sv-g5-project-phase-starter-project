import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_event.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_state.dart';
import 'package:ecommerce/features/product/presentation/screens/add_product_page.dart';
import 'package:ecommerce/injection_container.dart' as di;
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  setUp(() {
    di.init();
  });

  testWidgets('should create a product with valid inputs',
      (WidgetTester tester) async {
    // Arrange
    final mockProductBloc = MockProductBloc();
    when(mockProductBloc.state).thenReturn(ProductInitial());

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: BlocProvider<ProductBloc>(
          create: (_) => mockProductBloc,
          child: AddProductPage(),
        ),
      ),
    );

    await tester.enterText(find.byType(TextField).first, 'Test Product');
    await tester.enterText(find.byType(TextField).at(1), '10.0');
    await tester.enterText(
        find.byType(TextField).at(2), 'This is a test product.');

    await tester.tap(find.byType(ElevatedButton).first);
    await tester.pumpAndSettle();

    verify(mockProductBloc.add(InsertProductEvent(
      product: const Product(
        id: '',
        name: 'Test Product',
        price: 10.0,
        description: 'This is a test product.',
        imageUrl: '',
      ),
    )));
  });

  testWidgets('should not create a product with empty name',
      (WidgetTester tester) async {
    // Arrange
    final mockProductBloc = MockProductBloc();
    when(mockProductBloc.state).thenReturn(ProductInitial());

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: BlocProvider<ProductBloc>(
          create: (_) => mockProductBloc,
          child: const AddProductPage(),
        ),
      ),
    );

    await tester.enterText(find.byType(TextField).first, '');
    await tester.enterText(find.byType(TextField).at(1), '10.0');
    await tester.enterText(
        find.byType(TextField).at(2), 'This is a test product.');

    // Tap
    await tester.tap(find.byType(ElevatedButton).first);
    await tester.pumpAndSettle();

    // Assert
    verifyNever(mockProductBloc.add(InsertProductEvent(
      product: const Product(
        id: '',
        name: '',
        price: 10.0,
        description: 'This is a test product.',
        imageUrl: '',
      ),
    )));
  });


  
}
