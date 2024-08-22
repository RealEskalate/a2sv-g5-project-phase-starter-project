import 'dart:io';


import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/presentation/widgets/product_card.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {

  setUp((){
    HttpOverrides.global = null;
  });
  testWidgets('ProductCard widget test', (WidgetTester tester) async {
    
    final product = const Product(
      id: '1',
      imageUrl: 'https://example.com/image.png',
      name: 'Sample Product',
      price: 29.99,
      description: 'A sample product description.',
    );

    
    await tester.pumpWidget(MaterialApp(
      home: Scaffold(
        body: Padding(
          padding: const EdgeInsets.all(18.0),
          child: Column(
            children: [
              Expanded(child: ProductCard(product: product)),
            ],
          ),
        ),
      ),
      routes: {
        '/details': (context) => const Scaffold(
              body: Text('Navigation Successful'),
            ),
      },
    ));

    expect(find.byType(Image), findsOneWidget);
    
    expect(find.text('Sample Product'), findsOneWidget);
    expect(find.text('\$29.99'), findsOneWidget);

    
    await tester.tap(find.byType(ProductCard));
    await tester.pumpAndSettle();

    expect(find.text('Navigation Successful'), findsOneWidget);
  });
}
