import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:product_6/core/cubit/user_cubit.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_6/features/product/presentation/pages/add_update_page.dart';
import 'package:product_6/features/product/presentation/widgets/custom_textfiled.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

class MockUserCubit extends MockBloc<UserCubit, UserEntity?>
    implements UserCubit {}

Future<void> main() async {
  late MockProductBloc mockProductBloc;
  late MockUserCubit mockUserCubit;

  setUp(() {
    mockProductBloc = MockProductBloc();
    mockUserCubit = MockUserCubit();
    HttpOverrides.global = null;
  });

  tearDown(() {
    mockProductBloc.close();
    mockUserCubit.close();
  });

  testWidgets(
    'AddUpdatePage renders correctly',
    (WidgetTester tester) async {
      whenListen(
        mockProductBloc,
        Stream.fromIterable([
          AddProuctState(),
        ]),
        initialState: InitalState(),
      );

      await tester.pumpWidget(
        MaterialApp(
          home: BlocProvider<ProductBloc>.value(
            value: mockProductBloc,
            child: const AddUpdatePage(),
          ),
        ),
      );

      // Check that the AddUpdatePage widgets are present
      expect(find.text('Add Product'), findsOneWidget);
      expect(find.byType(CustomTextfiled), findsNWidgets(4));

      await tester.enterText(
          find.byType(CustomTextfiled).first, 'Sample Product');
      await tester.pump();

      // Verify button press
      await tester.tap(find.text('ADD'));
      await tester.pumpAndSettle();

      // Check if snack bar show message
      // expect(find.text('Invalid Price Input'), findsOneWidget);
    },
  );
}
