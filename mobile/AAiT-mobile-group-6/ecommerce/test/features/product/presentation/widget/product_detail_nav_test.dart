import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_state.dart';
import 'package:ecommerce/features/product/presentation/screens/homepage.dart';
import 'package:ecommerce/features/product/presentation/screens/product_detail_page.dart';
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

  const testProduct = Product(
    id: '1',
    name: 'Product 1',
    price: 10.0,
    description: 'Description 1',
    imageUrl: 'https://example.com/image1.png',
  );

  testWidgets(
    'tapping the back button on Product Detail page navigates to the homepage using Navigator.pushNamed',
    (WidgetTester tester) async {
      // Arrange: Set the initial state of the ProductBloc and provide it to the widget tree.
      when(mockProductBloc.state)
          .thenReturn(const AllProductsLoaded(products: [testProduct]));

      await tester.pumpWidget(
        MaterialApp(
          initialRoute: '/product_detail_page',
          routes: {
            '/homepage': (context) => BlocProvider<ProductBloc>(
                  create: (_) => mockProductBloc,
                  child: const HomePage(),
                ),
            '/product_detail_page': (context) => BlocProvider<ProductBloc>(
                  create: (_) => mockProductBloc,
                  child: const ProductDetailPage(),
                ),
          },
        ),
      );

      // Act: Tap the back button
      await tester.tap(find.byIcon(Icons.arrow_back_ios_new));
      await tester.pumpAndSettle();

      // Assert: Check that we are back to the homepage
      expect(find.byType(HomePage), findsOneWidget);
      expect(find.text('Product 1'), findsOneWidget);
    },
  );
}
