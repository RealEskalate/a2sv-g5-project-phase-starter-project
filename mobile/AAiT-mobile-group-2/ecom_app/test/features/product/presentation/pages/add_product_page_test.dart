import 'package:bloc_test/bloc_test.dart';
import 'package:ecom_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecom_app/features/product/presentation/pages/add_product_page.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends Mock implements ProductBloc {}

class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}

void main() {
  late MockProductBloc mockProductBloc;
  late MockAuthBloc mockAuthBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
    mockAuthBloc = MockAuthBloc();
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());
    when(() => mockAuthBloc.state).thenReturn(AuthInitial());
  });

  Future<void> _pumpWidget(WidgetTester tester) async {
    await tester.pumpWidget(
      MaterialApp(
        routes: {
          '/home': (context) => const Scaffold(
              body: Text('Navigation Successful'),
            ),
        },
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ProductBloc>.value(
              value: mockProductBloc,
            ),
            BlocProvider<AuthBloc>.value(
              value: mockAuthBloc,
            ),
          ],
          child: const AddProductPage(
            isAdd: true,
          ),
        ),
      ),
    );
  }

  const testProduct = Product(
    id: '1',
    name: 'Test Product',
    price: 100.0,
    description: 'Test Description',
    imageUrl: '',
  );

  Future<void> _pumpUpdateWidget(WidgetTester tester) async {
    await tester.pumpWidget(
      MaterialApp(
        routes: {
          '/home': (context) => const Scaffold(
              body: Text('Navigation Successful'),
            ),
        },
        home: BlocProvider<ProductBloc>.value(
          value: mockProductBloc,
          child: const AddProductPage(
            isAdd: false,
            product: testProduct,
          ),
        ),
      ),
    );
  }

  testWidgets('shows success Snackbar on successful product creation',
      (WidgetTester tester) async {
    // Arrange
    const mockProduct = Product(
      id: '1',
      name: 'Test Product',
      price: 50.0,
      description: 'Test Description',
      imageUrl: 'test_image_url',
    );
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable(
          [ProductCreatedState(product: mockProduct)]),
    );

    await _pumpWidget(tester);

    await tester.enterText(find.byType(TextField).at(0), 'Test Product');
    await tester.enterText(find.byType(TextField).at(1), 'Test Category');
    await tester.enterText(find.byType(TextField).at(2), '50');
    await tester.enterText(find.byType(TextField).at(3), 'Test Description');

    final addButton = find.text('ADD').first;
    // await tester.ensureVisible(addButton);
    await tester.tap(addButton, warnIfMissed: false);
    await tester.pumpAndSettle();

    expect(find.text('Product Created Successfully'), findsOneWidget);
  });

  testWidgets('shows error Snackbar on product creation failure',
      (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable(
          [ProductCreatedErrorState(message: 'Failed to create product')]),
    );

    await _pumpWidget(tester);

    await tester.enterText(find.byType(TextField).at(0), 'Test Product');
    await tester.enterText(find.byType(TextField).at(1), 'Test Category');
    await tester.enterText(find.byType(TextField).at(2), '50');
    await tester.enterText(find.byType(TextField).at(3), 'Test Description');

    await tester.tap(find.text('ADD'));
    await tester.pumpAndSettle();

    expect(find.text('Failed to create product'), findsOneWidget);
  });

  testWidgets('shows success Snackbar on successful product update',
      (WidgetTester tester) async {
    // Arrange
    const mockProduct = Product(
      id: '1',
      name: 'Test Product',
      price: 50.0,
      description: 'Test Description',
      imageUrl: 'test_image_url',
    );
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable(
          [ProductUpdatedState(product: mockProduct)]),
    );

    await _pumpUpdateWidget(tester);

    // Simulate filling out form fields
    await tester.enterText(find.byType(TextField).at(0), 'Updated Product');
    await tester.enterText(find.byType(TextField).at(1), 'Updated Category');
    await tester.enterText(find.byType(TextField).at(2), '200');
    await tester.enterText(find.byType(TextField).at(3), 'Updated Description');

    // Simulate tapping the UPDATE button
    await tester.tap(find.text('UPDATE').first, warnIfMissed: false);
    await tester.pumpAndSettle();

    expect(find.text('Product Updated Successfully'), findsOneWidget);
  });

  testWidgets('shows error Snackbar on product update failure',
      (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitialState());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable(
          [ProductUpdatedErrorState(message: 'Failed to update product')]),
    );

    await _pumpUpdateWidget(tester);

    await tester.enterText(find.byType(TextField).at(0), 'Updated Product');
    await tester.enterText(find.byType(TextField).at(1), 'Updated Category');
    await tester.enterText(find.byType(TextField).at(2), '200');
    await tester.enterText(find.byType(TextField).at(3), 'Updated Description');

    await tester.tap(find.text('UPDATE'));
    await tester.pumpAndSettle();

    expect(find.text('Failed to update product'), findsOneWidget);
  });
}
