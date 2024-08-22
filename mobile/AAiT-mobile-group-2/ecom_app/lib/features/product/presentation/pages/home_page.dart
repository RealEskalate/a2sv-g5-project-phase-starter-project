import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../auth/presentation/bloc/auth_bloc.dart';
import '../../domain/entities/product.dart';
import '../bloc/product_bloc.dart';
import '../widgets/product_card.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    context.read<AuthBloc>().add(GetUserEvent());
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is AuthLoggedOut) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Logged out successfully'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
          Navigator.pushNamedAndRemoveUntil(context, '/login', (Route<dynamic> route) => false,);
        }
      },
      child: Scaffold(
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            Navigator.of(context).pushNamed('/add');
          },
          child: const Icon(
            Icons.add,
            color: Colors.white,
          ),
          backgroundColor: Theme.of(context).primaryColor,
          shape: const CircleBorder(),
        ),
        appBar: AppBar(
          leadingWidth: 70,
          // toolbarHeight: 80,
          leading: Padding(
            padding: const EdgeInsets.only(left: 20),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(9),
              child: Image.asset(
                'assets/profile_picture.png',
                width: 45,
                height: 55,
              ),
            ),
          ),
          title: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Padding(
                padding: EdgeInsets.only(left: 5.0),
                child: Text(
                  'Aug 7, 2024',
                  style: TextStyle(
                      fontWeight: FontWeight.w300,
                      color: Colors.black45,
                      fontSize: 13),
                ),
              ),
              const SizedBox(
                height: 2,
              ),
              Row(
                children: [
                  const Text(
                    'Hello, ',
                    style: TextStyle(color: Colors.black45, fontSize: 15),
                  ),
                  BlocBuilder<AuthBloc, AuthState>(
                    builder: (context, state) {
                      if (state is AuthAuthenticated) {
                        return Text(
                          state.userDataEntity.name,
                          style: const TextStyle(
                              fontWeight: FontWeight.w700, fontSize: 15),
                        );
                      } else {
                        return const Text(
                          'User',
                          style: TextStyle(
                              fontWeight: FontWeight.w700, fontSize: 15),
                        );
                      }
                    },
                  )
                ],
              )
            ],
          ),
          actions: [
            Container(
                decoration: BoxDecoration(
                  border: Border.all(width: 2, color: Colors.grey.shade300),
                  borderRadius: BorderRadius.circular(5),
                ),
                child: InkWell(
                    onTap: () {},
                    splashColor: Colors.grey.shade300,
                    child: Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Stack(children: [
                        const Icon(Icons.notifications_none_rounded),
                        Positioned(
                            top: 3,
                            right: 5,
                            child: Container(
                              width: 6,
                              height: 6,
                              decoration: BoxDecoration(
                                  borderRadius: BorderRadius.circular(20),
                                  color:
                                      const Color.fromARGB(255, 63, 81, 243)),
                            ))
                      ]),
                    ))),
            const SizedBox(
              width: 15,
            ),
            IconButton(
                onPressed: () {
                  context.read<AuthBloc>().add(LogoutEvent());
                },
                icon: const Icon(Icons.logout))
          ],
        ),
        body: Padding(
            padding:
                const EdgeInsets.only(left: 15, right: 15, top: 30, bottom: 20),
            child: buildHome(context)),
      ),
    );
  }
}

Widget buildHome(BuildContext context) {
  context.read<ProductBloc>().add(LoadAllProductEvent());

  return Column(
    children: [
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const Text(
            'Available Products',
            style: TextStyle(fontWeight: FontWeight.w700, fontSize: 20),
          ),
          Container(
              decoration: BoxDecoration(
                border: Border.all(width: 2, color: Colors.grey.shade300),
                borderRadius: BorderRadius.circular(5),
              ),
              child: InkWell(
                  onTap: () {
                    Navigator.of(context).pushNamed('/search');
                  },
                  splashColor: Colors.grey.shade300,
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Icon(
                      Icons.search,
                      color: Colors.grey.shade600,
                    ),
                  )))
        ],
      ),
      const SizedBox(
        height: 10,
      ),
      BlocBuilder<ProductBloc, ProductState>(builder: (context, state) {
        if (state is ProductLoading) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else if (state is LoadAllProductState) {
          return Expanded(
            child: RefreshIndicator(
              onRefresh: () async {
                context.read<ProductBloc>().add(LoadAllProductEvent());
              },
              child: ListView.separated(
                  itemCount: state.products.length,
                  separatorBuilder: (BuildContext context, int index) {
                    return const SizedBox(
                      height: 10,
                    );
                  },
                  itemBuilder: (BuildContext context, int index) {
                    return ProductCard(
                      product: Product(
                          id: state.products[index].id,
                          imageUrl: state.products[index].imageUrl,
                          name: state.products[index].name,
                          price: state.products[index].price,
                          description: state.products[index].description),
                    );
                  }),
            ),
          );
        } else if (state is ProductErrorState) {
          return Center(
            child: Text('Error: ${state.message}'),
          );
        } else {
          return const Center(
            child: Text('No products available'),
          );
        }
      }),
    ],
  );
}
