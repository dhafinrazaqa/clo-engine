#!/usr/bin/env python3
"""
visuallize.py

Usage:
    python visualize.py output.json

Generates:
    - top_view.png (XY)
    - front_view.png (XZ)
    - side_view.png (YZ)

Requirements:
    pip install matplotlib
"""
import sys
import json
import math
import matplotlib.pyplot as plt
from matplotlib.patches import Rectangle

def load_json(path):
    with open(path, 'r', encoding='utf-8') as f:
        return json.load(f)

def draw_top_view(container, placements, outpath='top_view.png', scale=None):
    L = container['length']
    W = container['width']

    if scale is None:
        # scale to width 12 inches-ish in figure: determine scale factor so longest side ~12
        max_dim = max(L, W)
        scale = max_dim / 12.0 if max_dim > 0 else 1

    fig_w = max(6, L/scale/2)
    fig_h = max(4, W/scale/2)
    fig, ax = plt.subplots(figsize=(fig_w, fig_h))
    ax.set_title('Top view (X = length, Y = width)')
    ax.set_xlim(0, L)
    ax.set_ylim(0, W)
    ax.set_xlabel('X (length)')
    ax.set_ylabel('Y (width)')
    ax.set_aspect('equal', adjustable='box')
    ax.invert_yaxis() # top-left origin

    for p in placements:
        x = p['x']
        y = p['y']
        length = p['length']
        width = p['width']

        rect = Rectangle((x, y), length, width, fill=True, alpha=0.6, ec='k')
        ax.add_patch(rect)
        ax.text(x + length/2, y + width/2, p['item_instance_id'], ha='center', va='center', fontsize=8)

    ax.grid(True, linestyle='--', alpha=0.3)
    fig.tight_layout()
    fig.savefig(outpath, dpi=150)
    plt.close(fig)

def draw_front_view(container, placements, outpath='front_view.png', scale=None):
    # X (length) horizontally, Z (height) vertically
    L = container['length']
    H = container['height']

    fig_w = max(6, L/100.0)
    fig_h = max(4, H/100.0)
    fig, ax = plt.subplots(figsize=(fig_w, fig_h))
    ax.set_title('Front view (X = length, Z = height)')
    ax.set_xlim(0, L)
    ax.set_ylim(0, H)
    ax.set_xlabel('X (length)')
    ax.set_ylabel('Z (height)')
    ax.set_aspect('auto')

    for p in placements:
        x = p['x']
        z = p['z']
        length = p['length']
        height = p['height']

        rect = Rectangle((x, z), length, height, fill=True, alpha=0.6, ec='k')
        ax.add_patch(rect)
        ax.text(x + length/2, z + height/2, p['item_instance_id'], ha='center', va='center', fontsize=7)

    ax.grid(True, linestyle='--', alpha=0.3)
    fig.tight_layout()
    fig.savefig(outpath, dpi=150)
    plt.close(fig)

def draw_side_view(container, placements, outpath='side_view.png', scale=None):
    # Y (width) horizontally, Z (height) vertically
    W = container['width']
    H = container['height']

    fig_w = max(6, W/100.0)
    fig_h = max(4, H/100.0)
    fig, ax = plt.subplots(figsize=(fig_w, fig_h))
    ax.set_title('Side view (Y = width, Z = height)')
    ax.set_xlim(0, W)
    ax.set_ylim(0, H)
    ax.set_xlabel('Y (width)')
    ax.set_ylabel('Z (height)')
    ax.set_aspect('auto')

    for p in placements:
        y = p['y']
        z = p['z']
        width = p['width']
        height = p['height']

        rect = Rectangle((y, z), width, height, fill=True, alpha=0.6, ec='k')
        ax.add_patch(rect)
        ax.text(y + width/2, z + height/2, p['item_instance_id'], ha='center', va='center', fontsize=7)

    ax.grid(True, linestyle='--', alpha=0.3)
    fig.tight_layout()
    fig.savefig(outpath, dpi=150)
    plt.close(fig)

def normalize_placements(raw):
    # ensure required keys exists and conver strings/numbers consistently
    placements = []
    for p in raw:
        # some fields might be null in output, ensure ints
        pp = {
            'item_instance_id': p.get('item_instance_id'),
            'x': int(p.get('x', 0)),
            'y': int(p.get('y', 0)),
            'z': int(p.get('z', 0)),
            'length': int(p.get('length', 0)),
            'width': int(p.get('width', 0)),
            'height': int(p.get('height', 0)),
            'orientation': p.get('orientation'),
            'shelf_index': p.get('shelf_index'),
            'shelf_start_z': p.get('shelf_start_z'),
            'shelf_height': p.get('shelf_height'),
        }
        placements.append(pp)
    return placements

def main():
    if len(sys.argv) < 2:
        print("Usage: python visualize.py output.json")
        sys.exit(1)

    path = sys.argv[1]
    data = load_json(path)

    if 'placements' not in data or 'metrics' not in data:
        print("Invalid output JSON: missing placements or metrics")
        sys.exit(1)

    container = {
        'length': data.get('container', {}).get('length') if data.get('container') else None,
        'width': data.get('container', {}).get('width') if data.get('container') else None,
        'height': data.get('container', {}).get('height') if data.get('container') else None,
    }

    # fallback: infer container size from metrics if there is none
    if not container['length'] or not container['width'] or not container['height']:
        # attempt to read container from top-level metrics shape (best effort)
        print("Warning: container dimensions missing in JSON, trying default values")
        # (subject to change)
        container = {'length': 1200, 'width': 235, 'height': 269}

    raw_placements = data['placements']
    placements = normalize_placements(raw_placements)

    draw_top_view(container, placements, outpath="top_view.png")
    draw_front_view(container, placements, outpath="front_view.png")
    draw_side_view(container, placements, outpath="side_view.png")

    print("Saved: top_view.png, front_view.png, side_view.png")

if __name__ == '__main__':
    main()
