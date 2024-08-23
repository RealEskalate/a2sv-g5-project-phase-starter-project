import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_8/features/product/presentation/bloc/product_state.dart';
import 'package:product_8/features/product/presentation/pages/add_product_page.dart';

// Mock the bloc
class MockProductBloc extends Mock implements ProductBloc {}

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
  });

  Future<void> _pumpWidget(WidgetTester tester) async {
    await tester.pumpWidget(
      MaterialApp(
        home: BlocProvider<ProductBloc>.value(
          value: mockProductBloc,
          child: const ADDPage(),
        ),
      ),
    );
  }

  testWidgets('shows success Snackbar on successful product creation', (WidgetTester tester) async {
    // Arrange
    const  mockProduct =  Product(
      id: '1',
      name: 'Test Product',
      price: 50.0,
      description: 'Test Description',
      imageUrl: 'test_image_url',
    );
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable([const ProductCreatedState(product: mockProduct)]),
    );

    await _pumpWidget(tester);

    // Simulate filling out form fields
    await tester.enterText(find.byType(TextField).at(0), 'Test Product');
    await tester.enterText(find.byType(TextField).at(1), 'Test Category');
    await tester.enterText(find.byType(TextField).at(2), '50');
    await tester.enterText(find.byType(TextField).at(3), 'Test Description');

    // Simulate tapping the ADD button
    await tester.tap(find.text('ADD').first);
    await tester.pumpAndSettle();

    // Assert that the success Snackbar is shown
    expect(find.text('Product Created Successfully'), findsOneWidget);
  });

  testWidgets('shows error Snackbar on product creation failure', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable([const ProductCreatedErrorState(message: 'Creation Failed')]),
    );

    await _pumpWidget(tester);

    // Simulate filling out form fields
    await tester.enterText(find.byType(TextField).at(0), 'Test Product');
    await tester.enterText(find.byType(TextField).at(1), 'Test Category');
    await tester.enterText(find.byType(TextField).at(2), '50');
    await tester.enterText(find.byType(TextField).at(3), 'Test Description');

    // Simulate tapping the ADD button
    await tester.tap(find.text('ADD'));
    await tester.pumpAndSettle();

    // Assert that the error Snackbar is shown
    expect(find.text('Creation Failed'), findsOneWidget);
  });
  
}
