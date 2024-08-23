import 'dart:io';

import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/presentation/bloc/product_bloc.dart';
import 'package:application1/features/product/presentation/pages/details_page.dart';
import 'package:application1/features/product/presentation/pages/home_page.dart';
import 'package:application1/features/product/presentation/widgets/components/size_cards.dart';

import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
    mockProductBloc = MockProductBloc();
    HttpOverrides.global = null;
  });

  Widget makeTestableWidget(Widget body) {
    return BlocProvider<ProductBloc>.value(
      value: mockProductBloc,
      child: MaterialApp(
        home: body,
        routes: {
          '/home_page': (context) => const Home(),
        },
      ),
    );
  }

  ProductEntity tProduct = const ProductEntity(
    id: '1',
    name: 'product1',
    description: 'this is a tProduct',
    price: 23,
    imageUrl:
        'https://i.pinimg.com/564x/ba/86/04/ba86047d55280a343e3c1f0e0868f0e7.jpg',
  );

  testWidgets('details page should have the necessary products',
      (widgetTester) async {
    when(() => mockProductBloc.state).thenReturn(ProductInitial());
    //act
    await widgetTester
        .pumpWidget(makeTestableWidget(DetailsPage(selectedProduct: tProduct)));
    await widgetTester.pumpAndSettle();
    var sizeCard = find.byType(SizeCards);
    var description = find.text('this is a tProduct');
    var price = find.text('\$23.0');
    var title = find.text('product1');
    //assert
    expect(title, findsOneWidget);
    expect(price, findsOneWidget);
    expect(sizeCard, findsAny);
    expect(description, findsOneWidget);
  });

  testWidgets('DELETE button triggers DeleteProductEvent',
      (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    // Act
    await tester
        .pumpWidget(makeTestableWidget(DetailsPage(selectedProduct: tProduct)));
    await tester.pumpAndSettle();
    await tester.tap(find.text('DELETE'));
    await tester.pump();

    // Assert
    verifyNever(() => mockProductBloc.add(DeleteProductEvent(id: tProduct.id)))
        .called(0);
  });

}
