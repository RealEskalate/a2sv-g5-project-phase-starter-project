import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_8/features/product/presentation/bloc/product_state.dart';
import 'package:product_8/features/product/presentation/pages/update_page.dart';


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
          child: const UpDate(
            product: Product(
              id: '1',
              name: 'Test Product',
              price: 100.0,
              description: 'Test Description',
              imageUrl: '',
            ),
          ),
        ),
      ),
    );
  }

  testWidgets('shows success Snackbar on successful product update', (WidgetTester tester) async {
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
      Stream<ProductState>.fromIterable([const ProductUpdatedState(product: mockProduct)]),
    );

    await _pumpWidget(tester);

    // Simulate filling out form fields
    await tester.enterText(find.byType(TextField).at(0), 'Updated Product');
    await tester.enterText(find.byType(TextField).at(1), 'Updated Category');
    await tester.enterText(find.byType(TextField).at(2), '200');
    await tester.enterText(find.byType(TextField).at(3), 'Updated Description');

    // Simulate tapping the UPDATE button
    await tester.tap(find.text('UPDATE').first);
    await tester.pumpAndSettle();

    // Assert that the success Snackbar is shown
    expect(find.text('Product Updated Successfully'), findsOneWidget);
  });

  testWidgets('shows error Snackbar on product update failure', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    whenListen(
      mockProductBloc,
      Stream<ProductState>.fromIterable([const ProductUpdatedErrorState(message: 'Update Failed')]),
    );

    await _pumpWidget(tester);

    // Simulate filling out form fields
    await tester.enterText(find.byType(TextField).at(0), 'Updated Product');
    await tester.enterText(find.byType(TextField).at(1), 'Updated Category');
    await tester.enterText(find.byType(TextField).at(2), '200');
    await tester.enterText(find.byType(TextField).at(3), 'Updated Description');

    // Simulate tapping the UPDATE button
    await tester.tap(find.text('UPDATE'));
    await tester.pumpAndSettle();

    // Assert that the error Snackbar is shown
    expect(find.text('Update Failed'), findsOneWidget);
  });

 
}
